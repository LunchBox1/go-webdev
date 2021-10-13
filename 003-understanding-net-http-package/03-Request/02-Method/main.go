package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) { //handler
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values //
	}{
		req.Method, //request method (POST or GET); here we are asking for the method from our http request
		req.Form,   //type Values map[string][]string
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
