package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func router(req Request) func(Request) string {

	// static files
	staticRxp, _ := regexp.Compile("/static/(.*)")
	if len(staticRxp.FindStringSubmatch(req["path"])) > 0 {
		return func(req Request) string {
			return getStatic(staticRxp.FindStringSubmatch(req["path"])[1], req)
		}
	}

	// index route
	if req["path"] == "/" {
		return getIndex
	}

	// posts route
	if req["path"] == "/post" {
		return getPost
	}

	if req["path"] == "/favicon.ico" {
		return func(req Request) string {
			return getStatic("gopher3.png", req)
		}
	}

	// 404 if nothing else matched
	return notFound
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
	imagePath, _ := filepath.Abs("static/" + path)
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
