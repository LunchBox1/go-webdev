package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") //ParseFiles returns 2 things; a val of type *Template and val of type error. This is why we are able to declare two variables (tpl and err).
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil) //Execute is a method that has type *Template attached to it. Execute returns an error. Execute takes in a writter and data. Here we are passing in
	//                                  os.Stdout(a writter interface) and not data(nil). A writter is just a location to where the data is executed to.
	if err != nil {
		log.Fatalln(err)
	}
}
