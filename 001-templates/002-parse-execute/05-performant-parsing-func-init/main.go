package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template //tpl is of type *template from package Template.

func init() { //func init initializes the program. It runs once when the program is starting. In this case, it will parse the template all at once in the beggining.
	tpl = template.Must(template.ParseGlob("templates/*.gmao")) //template.ParseFiles returns type *Template. This makes tpl type *Template. ParseGlob parses a folder location. You can
	//limit to files that have a specific extension.
	//In this situation, tamplate.Must is doing the error checking. The Must function, from package template, take in a *Template and an error. These two things are provided by the
	//template.ParseGlob.
}

func main() {
	err := tpl.Execute(os.Stdout, nil) //Execute is a method attached to type *Template. Will execute the first template it finds.
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
