package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Open a TCP connection on a specified port
	service := ":1234"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("listening on: ", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		fmt.Println("Could not open TCP connection", err)
	}

	// close the tcp listener when the application closes
	defer listener.Close()

	// Main server loop
	// * listen for a connection
	// * generate and write responses
	// * close the connection
	for {
		conn, err := listener.Accept() // Blocking. Waits for connection
		if err != nil {
			fmt.Println("recieved connection, but errored", err)
			continue
		}

		// Handle incoming connections in a thread
		go func() {
			// build a Request from the connection (map of http method, path, and headers)
			request := MakeRequestFromConnection(conn)

			// Log the incoming request
			fmt.Printf("\033[1;36m%s\033[0m", time.Now().Format("2006-01-02 15:04:05")+" [Go Server] ")
			fmt.Printf("%+v\n", request)

			// Generate and write a response, and close the connection
			responseString := router(request)(request)

			conn.Write([]byte(responseString))
			conn.Close()
		}()
	}
}
