package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) { //any type of with this method is also of type http.Handler interface
	fmt.Println("Any code you want in this func")
}

func main() {

}
