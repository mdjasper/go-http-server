# go-http-server

This is an experimental/learning project with the goal of building an http server
"from scratch" starting at the TCP layer, without using any of the built-in
HTTP libraries provided by go.

## Program Execution

`main.go` opens a TCP connection and waits for incoming requests forever while
the program is running. Incomming requests are handled outside of the main thread
in go routines. The incoming request stream is read, and parsed into a Request
object. Then basic pattern matching is applied to the request, and it is routed
to one of several handlers (static, not found, index, etc). The handlers generate
an appropriate response (html, image asset, etc) and the response is written
to the TCP connection which is then closed.

## Routing

The router is a function which returns a function which returns the response
(headers + body) as a string. These handlers can be anonymous closures or named
functions, as long as the method signature matches. Path matches are evaluated
in the order that they are declared.

## Usage

go-http-server may be useful for study of basic HTTP server architecture, but has
not been evaluated for security, resource usage, etc. It would not be appropriate
to use go-http-server for any purpose beyond study or research.