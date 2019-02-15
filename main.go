package main

import (
  "fmt"
  "net"
  "os"
  "log"
  "io/ioutil"
  "path/filepath"
  "regexp"
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

    go func(){
      // Read connection into a byte array
      // and convert into a Request
      buffer := make([]byte, 1024)
      reqLen, err := conn.Read(buffer)
      if err != nil {
        fmt.Println("Error reading from connection", err)
      }
      request := MakeRequestFromString(string(buffer[:reqLen]))

      log.Printf("%+v\n", request)

      var responseBody string

      //default to 404
      responseBody = notFound(request)

      //look for static requests
      staticRxp, err := regexp.Compile("/static/(.*)")
      if len(staticRxp.FindStringSubmatch(request.path)) > 0 {
        responseBody = getStatic(staticRxp.FindStringSubmatch(request.path)[1], request)
      }

      //custom routes
      if request.path == "/" { responseBody = getIndex(request) }
      if request.path == "/post" { responseBody = getIndex(request) }

      conn.Write([]byte(responseBody))
      conn.Close()
    }()
  }
}

func getIndex(req Request) string {
  body := `<html><body>
<h1>home page</h1>
<p><a href="/post">post</a></p>
<p><img src="/static/golang_128x128.png"/></p>
</body></html>`

  headers := DefaultHtmlHeaders()

  return response(200, headers, body)
}

func getPost(req Request) string {
  body := `<html><body><h1>A Post</h1>
<p>Lorem Ipsom</p>
<p><a href="/">index</a></p></body></html>`

  headers := DefaultHtmlHeaders()

  return response(200, headers, body)
}

func getStatic(path string, req Request) string {
  imagePath, _ := filepath.Abs("static/"+ path)
  fmt.Println("finding ", imagePath)
  file, err := ioutil.ReadFile(imagePath)
  if err != nil {
    return notFound(req)
  }
  body := string(file[:])
  headers := PngHeaders()
  return response(200, headers, body)
}

func notFound(req Request) string {
  body := `<html><body><h1>404</h1>
<p>Page Not Found</p></body></html>`

  headers := DefaultHtmlHeaders()

  return response(200, headers, body)
}

func checkError(err error) {
  if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
  }
}