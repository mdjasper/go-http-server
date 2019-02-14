package main

import (
  "fmt"
  "net"
  "os"
  "strconv"
  "log"
)

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
    conn, err := listener.Accept() // Blocking. Waits for connection
    if err != nil {
      log.Println("recieved connection, but errored: ", err)
      continue
    }

    go func(){
      // Read connection into a byte array
      // and convert into a Request
      buffer := make([]byte, 1024)
      reqLen, err := conn.Read(buffer)
      if err != nil {
        fmt.Println("Error reading from connection: ", err)
      }
      request := MakeRequestFromString(string(buffer[:reqLen]))

      log.Printf("%+v\n", request)

      var responseBody string

      switch request.path {
      case "/":
        responseBody = getIndex(request)
      case "/post":
        responseBody = getPost(request)
      default:
        responseBody = notFound(request)
      }

      conn.Write([]byte(responseBody))
      conn.Close()
    }()
  }
}

func getIndex(req Request) string {
  body := `<h1>home page</h1>
<p><a href="/post">post</a></p>`

  headers := make(map[string]string)

  headers["Server"] = "JasperGo"
  headers["Content-Type"] = "text/html"
  headers["Connection"] = "Keep-Alive"
  headers["Keep-Alive"] = "timeout=5, max=997"
  headers["Transfer-Encoding"] = "identity"
  headers["Content-Length"] = strconv.Itoa(len(body))

  return response(200, headers, body)
}

func getPost(req Request) string {
  body := `<h1>A Post</h1>
<p>Lorem Ipsom</p>
<p><a href="/">index</a></p>`

  headers := `HTTP/1.1 200 OK
Server: JasperGo
Content-type: text/html
Connection: Keep-Alive
Keep-Alive: timeout=5, max=997
Transfer-Encoding: identity
Content-Length: ` + strconv.Itoa(len(body))

  return headers + "\r\n\r\n" + body
}

func notFound(req Request) string {
  body := `<h1>404</h1>
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

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}