package main

import (
	"strings"
)

type Request map[string]string

func MakeRequestFromString(requestString string) Request {
	//terms taken from https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html
	r := make(Request)
	lines := strings.Split(requestString, "\r\n")
	messageType, messageHeaders := lines[0], lines[1:]

	// parse messageType
	// example `GET /path HTTP/1.1`
	messageTypeParts := strings.Split(messageType, " ")
	r["path"] = messageTypeParts[1]
	r["method"] = messageTypeParts[0]

	// parse messageHeaders
	// example `Accept-Encoding: gzip`
	for _, line := range messageHeaders {
		lineParts := strings.Split(line, ": ")
		if len(lineParts[1]) > 1 {
			r[lineParts[0]] = lineParts[1]
		}
	}

	return r
}
