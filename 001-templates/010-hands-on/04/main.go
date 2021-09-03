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

type restaurant struct {
	Name string
	Menu []timeOfDay
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	Restaurants := []restaurant{
		restaurant{
			Name: "Boo Boo's",
			Menu: []timeOfDay{
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
			},
		},
		restaurant{
			Name: "Very Exclusive Dinner",
			Menu: []timeOfDay{
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
			},
		},
	}

	err := tpl.Execute(os.Stdout, Restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
