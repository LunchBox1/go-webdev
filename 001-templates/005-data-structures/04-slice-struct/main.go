package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}
	gandi := sage{
		Name:  "Gandi",
		Motto: "Be the change.",
	}
	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}
	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all.",
	}
	mohammad := sage{
		Name:  "Mohammad",
		Motto: "To overcome evil with goos is good, to resist evil with evil is evil.",
	}

	sages := []sage{buddha, gandi, mlk, jesus, mohammad}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
