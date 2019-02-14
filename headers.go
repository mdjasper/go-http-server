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