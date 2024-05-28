package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type tcpMessagePayload struct {
}

type tcpMessage struct {
	Module  string             `json:"module"`
	Type    string             `json:"type"`
	Payload *tcpMessagePayload `json:"payload"`
}

type tcpClient struct {
	conn *net.Conn
}

func (c *tcpClient) send(message *tcpMessage) {
	json.NewEncoder(*c.conn).Encode(message)
}

func (c *tcpClient) drop() {
	delete(tcpClients, c.conn)

	if err := (*c.conn).Close(); err != nil {
		log.Println("Error closing tcp connection:", err)
	}
}

var tcpClients = make(map[*net.Conn]*tcpClient)

func tcpServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5127")

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(tcpClients)
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
	//defer conn.Close()
	//defer tcpClients[&conn].drop()

	tcpClients[&conn] = &tcpClient{
		conn: &conn,
	}

loop:
	for {
		var msg tcpMessage
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

		fmt.Println("Received data:", msg)
	}

	tcpClients[&conn].drop()

	//for client := range

	//_, err = conn.Write(message)

}
