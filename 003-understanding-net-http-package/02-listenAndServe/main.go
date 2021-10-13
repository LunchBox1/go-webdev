package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) { //any type with this method is also of type http.Handler interface
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d) //func  ListenAndServe(addr string, handler Handler) error
}
