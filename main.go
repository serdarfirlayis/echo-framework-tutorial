package main

import (
	"fmt"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} */

type MyWebServerType bool

func (m MyWebServerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Hi
		</head>
		<body>
			<h1>Serdar Raşit</h1>
		</body>
	</html>
	`)

	/* fmt.Fprintln(w, "Hello World!")
	fmt.Fprintf(w, "Request is: %+v", r) */
}

func main() {
	var m MyWebServerType
	http.ListenAndServe("localhost:8080", m)
}
