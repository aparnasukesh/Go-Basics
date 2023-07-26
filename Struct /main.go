package main

import "fmt"

// Using type instead of var-----------------------------------------------------

type person struct {
	name  string
	age   int
	place string
}

// Using var---------------------------------------------------------------------

var Car struct {
	id    int
	name  string
	color string
}

// Another Example // Using type //----------------------------------------------

type menuItem struct {
	name   string
	prices map[string]float64
}

func main() {

	// Create an instance of the person struct
	person_1 := person{
		name:  "Aparna Sukesh",
		age:   23,
		place: "Kerala",
	}
	person_2 := person{
		name:  "Manjima",
		age:   23,
		place: "Kerala",
	}

	// Accessing fields
	fmt.Println(person_1)
	fmt.Println("Name:", person_1.name)
	fmt.Println("Age:", person_1.age)
	fmt.Println("place:", person_1.place)

	fmt.Println(person_2)
	fmt.Println("Name:", person_2.name)
	fmt.Println("Age:", person_2.age)
	fmt.Println("place:", person_2.place)

	// Creating instance of Car struct
	Car.id = 1
	Car.name = "BMW"
	Car.color = "Black"

	// Accessing fields
	fmt.Println(Car)
	fmt.Println(Car.id)
	fmt.Println(Car.name)

	// Creating instcance of menuItem struct
	menu := []menuItem{
		{
			name: "Coffee",
			prices: map[string]float64{
				"regular": 12.5,
				"medium":  14,
				"large":   15,
			},
		},
		{
			name: "Tea",
			prices: map[string]float64{
				"regular": 10,
				"medium":  20.5,
				"large":   25.5,
			},
		},
	}

	// Accessing fields
	fmt.Println(menu)
	fmt.Println(menu[0].name)
	fmt.Println(menu[0].prices)
	fmt.Println(menu[0].prices["regular"])

	fmt.Println(menu[1].name)
	fmt.Println(menu[1].prices)
	fmt.Println(menu[1].prices["regular"])

}
