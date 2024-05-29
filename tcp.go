package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"slices"
	"sync"
	"time"
)

type tcpMessagePayload struct {
}

type tcpMessage struct {
	Module  string `json:"module"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type tcpMessageRaw struct {
	//Id      string          `json:"id"`
	Module  string          `json:"module"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type tcpClient struct {
	conn *net.Conn
}

func (c *tcpClient) Send(message any) error {
	if err := json.NewEncoder(*c.conn).Encode(message); err != nil {
		return err
	}

	return nil
}

func (c *tcpClient) Drop() error {

	tcpClientsMutex.Lock()
	tcpEventSubsMutex.Lock()

	delete(tcpClients, c.conn)

	for event, clients := range tcpEventSubs {
		idx := slices.Index(clients, c)

		if idx == -1 {
			continue
		}

		tcpEventSubs[event] = slices.Delete(clients, idx, idx+1)
	}

	tcpEventSubsMutex.Unlock()
	tcpClientsMutex.Unlock()

	if err := (*c.conn).Close(); err != nil {
		return err
	}

	return nil
}

var tcpClientsMutex sync.Mutex
var tcpEventSubsMutex sync.Mutex
var tcpClients = make(map[*net.Conn]*tcpClient)
var tcpEventSubs = make(map[string][]*tcpClient) // [EventType]TcpClient

func tcpServer() {
	// Init event subscriptions variable
	for _, event := range events {
		tcpEventSubsMutex.Lock()
		tcpEventSubs[event] = make([]*tcpClient, 0)
		tcpEventSubsMutex.Unlock()
	}

	listener, err := net.Listen("tcp", "0.0.0.0:5127")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			//fmt.Println("TCP CLIENTS", tcpClients)
			//fmt.Println("TCP EVENT SUBS", tcpEventSubs)
		}
	}()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		log.Println("TCP client connected:", conn.RemoteAddr().String())

		go handleTcpConn(conn)
	}
}

func handleTcpConn(conn net.Conn) {
	tcpClientsMutex.Lock()

	tcpClients[&conn] = &tcpClient{
		conn: &conn,
	}

	tcpClientsMutex.Unlock()

loop:
	for {
		var msg tcpMessageRaw
		if err := json.NewDecoder(conn).Decode(&msg); err != nil {
			switch err {
			case io.EOF:
				log.Println("TCP client disconnected:", conn.RemoteAddr().String())
				break loop
			default:
				log.Println("Error decoding TCP message:", err)
				continue
			}
		}

		if !slices.Contains(events, msg.Type) {
			log.Println("Unknown type for subscribe:", msg.Type)
			continue
		}

		// @TODO for testing
		for _, v := range sseEventSubs[msg.Type] {
			err := v.Send(msg)
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}

		//fmt.Println("Received message type:", msg.Type)

		switch msg.Type {
		case "event/subscribe":
			log.Println("SUBSCRIBE")
			var payload []string
			if err := json.Unmarshal(msg.Payload, &payload); err != nil {
				log.Println("Error decoding payload:", err.Error())
				continue
			}

			for _, event := range payload {
				if !slices.Contains(events, event) {
					fmt.Println("UNKNOWN EVENT", event)
					continue
				}

				tcpEventSubsMutex.Lock()
				if !slices.Contains(tcpEventSubs[event], tcpClients[&conn]) {
					tcpEventSubs[event] = append(tcpEventSubs[event], tcpClients[&conn])
				}
				tcpEventSubsMutex.Unlock()
			}

		case "event/unsubscribe":
			var payload []string
			json.Unmarshal(msg.Payload, &payload)

			for _, event := range payload {
				if !slices.Contains(events, event) {
					log.Println("UNKNOWN EVENT", event)
					continue
				}

				tcpEventSubsMutex.Lock()
				idx := slices.Index(tcpEventSubs[event], tcpClients[&conn])

				if idx == -1 {
					continue
				}

				tcpEventSubs[event] = slices.Delete(tcpEventSubs[event], idx, idx+1)
				tcpEventSubsMutex.Unlock()
			}
		}

		tcpEventSubsMutex.Lock()
		for _, client := range tcpEventSubs[msg.Type] {
			if err := client.Send(msg); err != nil {
				log.Println("Error sending message:", err)
				client.Drop()
				return
			}
		}
		tcpEventSubsMutex.Unlock()

		//fmt.Println("Received data:", msg)
	}

	if err := tcpClients[&conn].Drop(); err != nil {
		log.Println("Error closing tcp connection:", err)
	}
}
