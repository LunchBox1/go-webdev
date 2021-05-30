package main

import "fmt"

/*number 1
func main() {
	x := []int{1, 3, 5, 7, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
}*/

/*number 2
func main() {
	m := map[string]int{
		"nik":   26,
		"manda": 27,
	}
	fmt.Println(m)

	for key := range m {
		fmt.Println(key)
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}*/

/*number 3, 4, 5, 6
type person struct {
	fname   string
	lname   string
	favfood []string
}

func (p person) walk() string {
	s := fmt.Sprintln(p.fname, "is walking")
	return s
}

func main() {
	p1 := person{
		fname:   "Danny",
		lname:   "Brown",
		favfood: []string{"pizza", "sushi", "borcht"},
	}
	fmt.Println(p1)
	fmt.Println(p1.fname)
	fmt.Println(p1.favfood)
	for _, v := range p1.favfood {
		fmt.Println(v)
	}
	fmt.Println(p1.walk())
}*/

/*number 7, 8, 9*/
type vehicle struct {
	doors int
	color string
}
type sadan struct {
	vehicle
	luxury bool
}
type truck struct {
	vehicle
	fourwheel bool
}
type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}
func (s sadan) transportationDevice() string {
	return "The ultimate driving experience."
}
func (t truck) transportationDevice() string {
	return "The ultimate off-road experience."
}
func main() {
	m2 := sadan{
		vehicle{
			doors: 2,
			color: "white",
		},
		true,
	}
	tacoma := truck{
		vehicle{
			doors: 4,
			color: "gray",
		},
		true,
	}
	fmt.Println(m2)
	fmt.Println(tacoma)
	fmt.Println(m2.vehicle.color)
	fmt.Println(tacoma.fourwheel)
	fmt.Println(m2.transportationDevice())
	fmt.Println(tacoma.transportationDevice())
	report(m2)
	report(tacoma)
}
