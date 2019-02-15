package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("listening on: ", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("Could not open TCP connection", err)
	}

	// close the tcp listener when that application closes
	defer listener.Close()

	for {
		conn, err := listener.Accept() // Blocking. Waits for connection
		if err != nil {
			fmt.Println("recieved connection, but errored", err)
			continue
		}

		go func() {
			// Read connection into a byte array
			// and convert into a Request
			buffer := make([]byte, 1024)
			reqLen, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading from connection", err)
			}

			//
			requestString := strings.TrimSpace(string(buffer[:reqLen]))
			request := MakeRequestFromString(requestString)
			fmt.Printf("\033[1;36m%s\033[0m", time.Now().Format("2006-01-02 15:04:05")+" [Go Server] ")
			fmt.Printf("%+v\n", request)

			responseString := router(request)(request)

			conn.Write([]byte(responseString))
			conn.Close()
		}()
	}
}
