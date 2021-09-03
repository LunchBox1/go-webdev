package main

import (
	"log"
	"os"
	"text/template"
)

type food struct {
	Name     string
	Calories string
}

type timeOfDay struct {
	TimeOfDay string
	Foods     []food
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	Menu := []timeOfDay{
		timeOfDay{
			TimeOfDay: "Breakfast",
			Foods: []food{
				food{
					Name:     "Pancakes",
					Calories: "500",
				},
				food{
					Name:     "Waffles",
					Calories: "600",
				},
			},
		},
		timeOfDay{
			TimeOfDay: "Lunch",
			Foods: []food{
				food{
					Name:     "Burger",
					Calories: "1000",
				},
				food{
					Name:     "Hotdog",
					Calories: "800",
				},
			},
		},
		timeOfDay{
			TimeOfDay: "Dinner",
			Foods: []food{
				food{
					Name:     "Steak",
					Calories: "1225",
				},
				food{
					Name:     "Salad",
					Calories: "425",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, Menu)
	if err != nil {
		log.Fatalln(err)
	}
}
