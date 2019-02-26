package main

import (
	"strings"
	"bufio"
	"net"
)

type Request map[string]string

func MakeRequestFromConnection(conn net.Conn) Request {

	// read connection stream line-by-line and parse into a Request
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html
	scanner := bufio.NewScanner(conn)
	r := make(Request)
	isMessageType := true

	for scanner.Scan() {
		line := scanner.Text()

		// stop scanning if we reached the end of the request body
		if line == "" {
			break
		}

		if isMessageType {
			// parse messageType
			// example `GET /path HTTP/1.1`
			messageTypeParts := strings.Split(line, " ")
			r["path"] = messageTypeParts[1]
			r["method"] = messageTypeParts[0]
			isMessageType = false
		} else {
			// parse messageHeaders
			// example `Accept-Encoding: gzip`
			lineParts := strings.Split(line, ": ")
			if len(lineParts[1]) > 1 {
				r[lineParts[0]] = lineParts[1]
			}
		}
	}

	return r
}
