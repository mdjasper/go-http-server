package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("listening on: ", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Println("Could not open TCP connection", err)
	}

	// close the tcp listener when that application closes
	defer listener.Close()

	for {
		conn, err := listener.Accept() // Blocking. Waits for connection
		if err != nil {
			log.Println("recieved connection, but errored", err)
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
			request := MakeRequestFromString(string(buffer[:reqLen]))

			log.Printf("%+v\n", request)

			resolveFn := router(request)
			responseString := resolveFn(request)

			conn.Write([]byte(responseString))
			conn.Close()
		}()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
