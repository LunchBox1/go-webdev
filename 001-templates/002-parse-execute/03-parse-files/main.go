package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	tpl, err := template.ParseFiles("one.gmao") //template.ParseFiles returns type *Template. This makes tpl type *Template.
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //Execute is a method attached to type *Template. Will execute the first template it finds.
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao") //passing in new templates to *Template. Using the ParseFiles method that is attached to type *Template
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil) //ExecuteTemplate is a method (attached to tpl) that executes a specific template.
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) //executed the first template it found.
	if err != nil {
		log.Fatalln(err)
	}
}
