package main

import "fmt"

type person struct {
	fname string
	lname string
}
type secretAgent struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println("Hello, my name is ", p.lname)
}
func (sa secretAgent) saSpeak() {
	fmt.Println("Shaken, not stirred")
}
func main() {
	p := person{
		fname: "James",
		lname: "Bond",
	}
	sa := secretAgent{
		person: p,
		ltk:    true,
	}
	fmt.Println(p.fname)
	p.pSpeak()
	fmt.Println(sa.lname)
	sa.saSpeak()
	sa.pSpeak()
}
