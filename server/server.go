package main

import (
	"fmt"
	"io"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	message := "Hello, Client!"
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err)
			}
			return
		}
		fmt.Println("Received message:", string(buffer))
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 50001")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleConnection(conn)
	}
}
