package main

import (
	"fmt"
	"strings"
)

type Headers map[string]string

func (h Headers) String() string {
	// strings.Build minimizes memory copying for each header value
	// which would occur with string concatenation
	var headerString strings.Builder
	for k, v := range h {
		fmt.Fprintf(&headerString, "%s: %s\r\n", k, v)
	}
	return headerString.String()
}

func DefaultHtmlHeaders() Headers {
	headers := make(Headers)

	headers["Server"] = "go-http-server"
	headers["Content-Type"] = "text/html"
	headers["Connection"] = "Keep-Alive"
	headers["Keep-Alive"] = "timeout=5, max=997"
	headers["Transfer-Encoding"] = "identity"

	return headers
}

func PngHeaders() Headers {
	headers := make(Headers)

	headers["Server"] = "go-http-server"
	headers["Content-Type"] = "image/png"
	headers["Connection"] = "Keep-Alive"
	headers["Keep-Alive"] = "timeout=5, max=997"
	headers["Transfer-Encoding"] = "identity"
	headers["Content-Transfer-Encoding"] = "binary"
	headers["Content-Disposition"] = "inline;"

	return headers
}
