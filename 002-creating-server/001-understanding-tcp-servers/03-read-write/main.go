package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn) //func NewScanner(r io.Reader) *Scanner; NewScanner returns a new Scanner to read from r.
	for scanner.Scan() {              //func (s *Scanner) Scan() bool; Scan advances the Scanner to the next token, which will then be available through the Bytes or Text method. It returns false when the scan stops,
		//either by reaching the end of the input or an error. scanner.Scan will scan until it return false. It will return false when it is done readin from the connection. The loop exits when scanner returns
		//false.
		ln := scanner.Text() //func (s *Scanner) Text() string; Text return the most recetn token generated by a call to Scan as a newly allocated string holding its bytes.
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) //Fprintf formats according to a forma spacifier and writes to w. It returns the number of butes written and any write error encountered.
		//func Fprintf(w io.Writer, format string, a ...interface{})(n int, err error)
	}
	defer conn.Close()

	//we never get here
	//we have an opem stream connection
	//how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
