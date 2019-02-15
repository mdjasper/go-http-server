package main

import(
  "strings"
)

type Request struct {
  host string
  userAgent string
  accept string
  method string
  path string
}

func MakeRequestFromString(requestString string) Request{
  r := Request{}
  lines := strings.Split(requestString, "\r\n")
  header, fieldsBody := lines[0], lines[1:]
  headerFields := strings.Split(header, " ")
  r.path = headerFields[1]
  r.method = headerFields[0]
  for _, line := range fieldsBody{
    kv := strings.Split(line, ": ")
    switch kv[0] {
    case "Host":
      r.host = kv[1]
    case "User-Agent":
      r.userAgent = kv[1]
    case "Accept":
      r.accept = kv[1]
    }
  }
  return r
}