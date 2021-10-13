package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

//this code programs a client that connects to a the TCP server.
//this one reads from the server.

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
