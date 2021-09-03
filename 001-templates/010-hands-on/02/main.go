package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}
type region struct {
	Region string
	Hotels []hotel
}
type regions struct {
	Southern, Central, Northern region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	California := regions{
		Southern: region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "San Diego Inn",
					Address: "2341 La Jolla Dr.",
					City:    "San Diego",
					Zip:     "92316",
				},
				hotel{
					Name:    "Hotel Cukamonga",
					Address: "0292 Foothill blwd.",
					City:    "Rancho Cukamonga",
					Zip:     "92390",
				},
			},
		},
		Northern: region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "The San Fran",
					Address: "123 Golden State Dr.",
					City:    "San Francisco",
					Zip:     "90845",
				},
				hotel{
					Name:    "Sons of Sleep",
					Address: "Charming Blvd.",
					City:    "Stockton",
					Zip:     "03099",
				},
			},
		},
		Central: region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Mammoth Inn",
					Address: "3432 Mammoth Dr.",
					City:    "Mammoth Lakes",
					Zip:     "92301",
				},
				hotel{
					Name:    "Ocean Stranded",
					Address: "Ocean Dr.",
					City:    "Santa Barbara",
					Zip:     "00435",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, California)
	if err != nil {
		log.Fatalln(err)
	}

}
