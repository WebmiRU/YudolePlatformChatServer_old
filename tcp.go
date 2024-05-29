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
	Event   string `json:"event"`
	Payload any    `json:"payload"`
}

type tcpMessageRaw struct {
	//Id      string          `json:"id"`
	Module  string          `json:"module"`
	Event   string          `json:"event"`
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

func (c *tcpClient) drop() error {
	delete(tcpClients, c.conn)

	if err := (*c.conn).Close(); err != nil {
		return err
	}

	return nil
}

var tcpClientsMutex sync.Mutex
var tcpClients = make(map[*net.Conn]*tcpClient)

func tcpServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5127")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("TCP CLIENTS", tcpClients)
			fmt.Println("EVENT SUBS", eventSubs)

			tcpClientsMutex.Lock()
			for _, v := range tcpClients {
				if err := v.Send(&tcpMessage{
					Module:  "example4",
					Event:   "message/chat",
					Payload: &tcpMessagePayload{},
				}); err != nil {
					log.Println("Error sending TCP client message:", err)
					continue
				}
			}
			tcpClientsMutex.Unlock()
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

		if !slices.Contains(events, msg.Event) {
			log.Println("Unknown event for subscribe:", msg.Event)
			continue
		}

		fmt.Println("Received message event:", msg.Event)

		switch msg.Event {
		case "event/subscribe":
			var payload []string
			json.Unmarshal(msg.Payload, &payload)

			for _, event := range payload {
				if _, ok := eventSubs[event]; !ok {
					fmt.Println("UNKNOWN EVENT", event)
					continue
				}

				eventSubsMutex.Lock()
				eventSubs[event] = append(eventSubs[event], tcpClients[&conn])
				eventSubsMutex.Unlock()
			}

		case "event/unsubscribe":
			var payload []string
			json.Unmarshal(msg.Payload, &payload)

			for _, event := range payload {
				if _, ok := eventSubs[event]; !ok {
					fmt.Println("UNKNOWN EVENT", event)
					continue
				}

				eventSubsMutex.Lock()
				delIdx := slices.Index(eventSubs[event], tcpClients[&conn])

				if delIdx >= 0 {
					eventSubs[event] = slices.Delete(eventSubs[event], delIdx, delIdx+1)
				}

				eventSubsMutex.Unlock()
			}

		default:
			log.Println("Unknown message type:", msg.Event)
		}
		fmt.Println("Received data:", msg)
	}

	if err := tcpClients[&conn].drop(); err != nil {
		log.Println("Error closing tcp connection:", err)
	}
}
