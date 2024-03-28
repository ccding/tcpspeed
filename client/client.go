package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	N := 100
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error parsing argument:", err)
			return
		}
		N = n
	}

	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	message := "Hello, Server!"
	buffer := make([]byte, 1024)

	totalStartTime := time.Now()
	for i := 0; i < N; i++ {
		startTime := time.Now()
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		latency := time.Since(startTime)
		fmt.Printf("Round-trip latency: %d us\n", latency.Microseconds())
	}
	totalLatency := time.Since(totalStartTime)
	fmt.Printf("Average round-trip latency: %d us\n", totalLatency.Microseconds()/int64(N))
}
