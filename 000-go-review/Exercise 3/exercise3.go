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
type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hello, my name is ", p.lname)
}
func (sa secretAgent) speak() {
	fmt.Println("Shaken, not stirred")
}
func spy(h human) {
	fmt.Println(h)
	switch v := h.(type) {
	case person:
		fmt.Println(v.fname)
	case secretAgent:
		fmt.Println(v.lname)
	default:
		fmt.Println("unknown")
	}
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
	spy(sa)
	spy(p)
}
