package main

import (
  "fmt"
  "net"
  "os"
  "io/ioutil"
  "bytes"
  "strings"
)

type Request struct {
  host string
  userAgent string
  accept string
  method string
  path string
}

func main() {
  service := ":1200"
  tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkError(err)
  for {
      conn, err := listener.Accept()
      if err != nil {
        continue
      }

      request, error := readConnection(conn)
      if error != nil {
        continue
      }

      buf := new(bytes.Buffer)
      buf.ReadFrom(bytes.NewReader(request))
      s := buf.String()
      req := parseRequest(s)
      go handleClient(conn, req)
  }
}

func parseRequest(requestString string) (Request) {
  lines := strings.Split(requestString, "\r\n")
  header, fieldsBody := lines[0], lines[1:]
  headerFields := strings.Split(header, " ")
  req := Request{
    method: headerFields[0],
    path: headerFields[1],
  }
  for _, line := range fieldsBody{
    kv := strings.Split(line, ": ")
    switch kv[0] {
    case "Host":
      req.host = kv[1]
    case "User-Agent":
      req.userAgent = kv[1]
    case "Accept":
      req.accept = kv[1]
    }
  }

  return req
}
func readConnection(conn net.Conn) ([]byte, error) {
  fmt.Println("XXX about to read from connection")
  defer fmt.Println("XXX finished reading from connection")

  bytes, err := ioutil.ReadAll(conn)
  if err != nil {
    fmt.Println(err)
  }
  return bytes, err
}

func handleClient(conn net.Conn, req Request) {
  defer conn.Close()

  fmt.Printf("%+v\n", req)

  if req.path == "/pages" {
    getPages(req, conn)
  }
}

func getPages(req Request, conn net.Conn) {
  fmt.Println("in getPages")
  html := "<h1>Hello World</h1>"
  conn.Write([]byte(html))
}

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}