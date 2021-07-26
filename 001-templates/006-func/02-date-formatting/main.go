package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006") //01-02-2006 specifies the format we want to use to format our date. It will not return the literal numbers. Format is a method attached to type Time.
}

var fm = template.FuncMap{ //FuncMap is a type, not a func
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
