package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path { //*Request is a struct (underlying). URL is a type(underlying struct) that is an elemect withing the *Reqest struct. Path is a string field inside the URL struct.
	case "/dog":
		io.WriteString(res, "doggy doggy doggy")
	case "/cat":
		io.WriteString(res, "kitty kitty kitty")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
