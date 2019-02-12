package main

import (
  "fmt"
  "net"
  "os"
  "strings"
  "strconv"
  "log"
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
  fmt.Println("listening on: ", tcpAddr)
  checkError(err)
  listener, err := net.ListenTCP("tcp", tcpAddr)

  // close the tcp listener when that application closes
  defer listener.Close()

  checkError(err)

  for {
      conn, err := listener.Accept()
      if err != nil {
        continue
      }

      // Read connection into a byte array
      // and convert into a Request
      buffer := make([]byte, 1024)
      reqLen, err := conn.Read(buffer)
      if err != nil {
        fmt.Println("Error reading from connection: ", err)
      }
      request := parseRequest(string(buffer[:reqLen]))

      log.Println(request)



      routes := map[string]func(Request)(string){
        "/": getIndex,
        "post": getPost,
        "*": notFound,
      }

      conn.Write([]byte(routes[request.path](request)))


      // switch request.path {
      // case "/":
      //   conn.Write([]byte(getIndex(request)))
      // case "/post":
      //   conn.Write([]byte(getPost(request)))
      // default:
      //   conn.Write([]byte(notFound(request)))
      // }


      // conn.Write([]byte(response))
      conn.Close()
  }
}

func getIndex(req Request) (string) {
  body := `<style type="text/css">body{background: #222; color:#eaeaea;}</style>
<h1>home page<h1>`

  headers := `HTTP/1.1 200 OK
Server: JasperGo
Content-type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=997
Transfer-Encoding: identity
Content-Length: ` + strconv.Itoa(len(body))

  return headers + "\r\n\r\n" + body
}

func getPost(req Request) (string) {
  body := `<style type="text/css">body{background: #222; color:#eaeaea;}</style>
<h1>A Post<h1>
<p>Lorem Ipsom</p>`

  headers := `HTTP/1.1 200 OK
Server: JasperGo
Content-type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=997
Transfer-Encoding: identity
Content-Length: ` + strconv.Itoa(len(body))

  return headers + "\r\n\r\n" + body
}

func notFound(req Request) (string) {
  body := `<style type="text/css">body{background: #222; color:#eaeaea;}</style>
<h1>404<h1>
<p>Page Not Found</p>`

  headers := `HTTP/1.1 404 Not Found
Server: JasperGo
Content-type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=997
Transfer-Encoding: identity
Content-Length: ` + strconv.Itoa(len(body))

  return headers + "\r\n\r\n" + body
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

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}