package main

import ()

type Headers map[string]string

func (h Headers) String() string {
  headerString := ""
  for k,v := range h {
    headerString += k + ": " + v + "\r\n"
  }
  return headerString
}

func DefaultHtmlHeaders() Headers {
  headers := make(map[string]string)

  headers["Server"] = "JasperGo"
  headers["Content-Type"] = "text/html"
  headers["Connection"] = "Keep-Alive"
  headers["Keep-Alive"] = "timeout=5, max=997"
  headers["Transfer-Encoding"] = "identity"

  return headers
}

func PngHeaders() Headers {
  headers := make(map[string]string)

  headers["Server"] = "JasperGo"
  headers["Content-Type"] = "image/png"
  headers["Connection"] = "Keep-Alive"
  headers["Keep-Alive"] = "timeout=5, max=997"
  headers["Transfer-Encoding"] = "identity"
  headers["Content-Transfer-Encoding"] = "binary"
  headers["Content-Disposition"] = "inline;"

  return headers
}