package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type tcpMessagePayload struct {
}

type tcpMessage struct {
	Type    string            `json:"type"`
	Payload tcpMessagePayload `json:"payload"`
}

type tcpClient struct {
	conn net.Conn
}

var tcp map[*net.Conn]*tcpClient

func tcpServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5127")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("TCP client connected:", conn.RemoteAddr().String())

		go handleTcpConn(conn)
	}
}

func handleTcpConn(conn net.Conn) {
	defer conn.Close()

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
	//_, err = conn.Write(message)

}
