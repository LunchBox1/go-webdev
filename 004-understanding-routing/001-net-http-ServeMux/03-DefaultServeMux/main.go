package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	var d dog
	var c cat

	http.Handle("/dog/", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) //When the handler in ListenAndServe in nil, the DefaultServerMux is invoked.
}
