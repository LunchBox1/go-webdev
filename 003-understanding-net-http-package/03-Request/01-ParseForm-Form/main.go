package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) { //this polymorphs type hotdog to also being type Handler interface. type Handler interface { ServeHTTP(ResponseWriter, *Request) }
	err := req.ParseForm() //func (r *Request) ParseForm() error; *Request is a struct
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(res, "index.gohtml", req.Form) //Form is an element inside struct form
	//in order to use From, we need to call ParseForm first.
	//we need to parse the form before we can use it.
	//Form is of type url.Values.
	//url.Values = type Values map[string][]string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog //also of type Handler
	http.ListenAndServe(":8080", d)
}
