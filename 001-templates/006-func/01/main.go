package main

// a func can be a type.
// if you want to pass in a func into a template, you will need to use the Funcs methods from the template package. Funcs method is attached to the *Template type. Funcs method takes
//in a value of FuncMap an returns a value of *Template.

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml")) /*funcs must be added before a template can be parsed. template.New give us *Template type. *Template
	type is needed in order to use the Funcs method. Funcs method takes in a value of type FuncMap. FuncMap is needed to pass funcs into a template. In this example we are passing
	func ToUpper and firstThree into the tpl.gohtml template.
	template.New is a func that returns a value of type *Template. Funcs() is a method that is attached to type *Template. That is why we are able to "chain" (connect with ".") the
	above.
	*/
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
