package main

func response(status int, headers Headers, body string) string {
  return "HTTP/1.1 " +
          HttpStatusCodes[status] +
          "\r\n" + headers.String() +
          "\r\n\r\n" +
          body
}