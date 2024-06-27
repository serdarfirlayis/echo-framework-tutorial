package main

import (
	"fmt"
	"net/http"
)

type login int
type welcome int

func (l login) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Using GET")
	case "POST":

		fmt.Fprintln(w, "Using POST")
	}

	fmt.Fprintf(w, "on login page")
}

func (wl welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "on welcome page")
}

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
	// http.HandleFunc("/login", mylogin)
	// http.HandleFunc("/welcome/", mywelcome)
	// http.Handle("/login", http.HandlerFunc(mylogin))
	// http.Handle("/welcome/", http.HandlerFunc(mywelcome))
	var i login
	var j welcome
	http.Handle("/login", i)
	http.Handle("/welcome/", j)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
