package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Hello World</title>
    </head>
    <body>
        Hello World!
    </body>
</html>`,
	)
}
func good(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Good night</title>
    </head>
    <body>
        Good night!!!
    </body>
</html>`,
	)
}
func goAway(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
    <head>
        <title>Good night</title>
    </head>
    <body>
        Go away !!! Faster, little mouser!!!
    </body>
</html>`,
	)
}
func testServer2() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/Good bay", good)
	http.HandleFunc("/", goAway)
	http.ListenAndServe(":9098", nil)
}
