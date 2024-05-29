package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
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
					Type:    "message/chat",
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
		var msg tcpMessage
		//if err := json.NewDecoder(conn).Decode(&msg); err != nil {
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

		// @TODO refactoring needed, draft/demo
		switch msg.Type {
		case "subscribe/event":
			//fmt.Println("subscribe/event")
			fmt.Println("subscribe/event:", msg.Payload)

			// WORK
			var bf bytes.Buffer
			err := json.NewEncoder(&bf).Encode(msg.Payload)
			if err != nil {
				log.Println("Error encoding payload:", err)
			}

			fmt.Println("BYTES:", bf.Bytes())
			fmt.Printf("BYTES STRING: %s\n", bf.Bytes())
			var v1 []string
			err1 := json.NewDecoder(&bf).Decode(&v1)
			if err1 != nil {
				fmt.Println("Error decoding payload:", err1)
				continue
			}

			for _, v := range v1 {
				fmt.Println("SUB STRING:", v)
			}

		case "message/chat":
			fmt.Println("message/chat:", msg.Payload)
		default:
			log.Println("Unknown message type:", msg.Type)
		}
		fmt.Println("Received data:", msg)
	}

	if err := tcpClients[&conn].drop(); err != nil {
		log.Println("Error closing tcp connection:", err)
	}
}
