package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//This creates a TCP server using the net package.

func main() {
	li, err := net.Listen("tcp", ":8080") //func Listen takes in two strings; one string identifies the type of network the server is listening on, and the second string specifies the port.
	//net.Listen return two values; one is of type Listener(an interface) and the other is of type error. The Listener interface has methods Accept() (Conn, error), Close() error, and Addr() Addr attached
	//to it.
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() //Accept() is a method that is in the Listener interface. This loop will run forever, listening for connections. Accept() returns two values; one value is of type Conn,
		//the other is of type error. Type Conn is an interface that has both the Read(b []byte) (n int, err error) and Write(b []byte) (n into, err error). These two methods implement type Reader and
		//type Writter (Interfaces). In other words, Conn is both a Reader and a Writter and can be passed into functions that require a reader or a writter.
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP server\n") // io.WriteString asks for a Writer and a string. conn implemets the Writer interface and thus is a Writer.
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()

	}
}
