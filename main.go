package main

import (
	"fmt"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} */

func mylogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Login
		</head>
		<body>
			<h1>Please enter your username and password</h1>
		</body>
	</html>
	`)
}

func mywelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<html>
		<head>
			Hi
		</head>
		<body>
			<h1>Welcome.</h1>
		</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/login", mylogin)
	http.HandleFunc("/welcome/", mywelcome)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
