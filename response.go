package main

import (
	"strconv"
)

func response(status int, headers Headers, body string) string {
	headers["Content-Length"] = strconv.Itoa(len(body))

	return "HTTP/1.1 " +
		HttpStatusCodes[status] +
		"\r\n" + headers.String() +
		"\r\n" +
		body
}
