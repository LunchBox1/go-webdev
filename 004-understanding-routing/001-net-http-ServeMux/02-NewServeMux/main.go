package main

import (
	"io"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog, dog, dog")
}

type cat int

func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat, cat, cat")
}

func main() {
	var d dog
	var c cat

	mux := http.NewServeMux() //func NewServeMux() *ServeMux; type ServeMux is a struct. ServeMux is an HTTP request multiplexer.
	mux.Handle("/dog/", d)    //Handle is a method that is attached to type *ServeMux.
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux) //*ServeMux is also of type Handler interface{} because it has the ServeHTTP(ResponseWriter, *Request) method attached.
}
