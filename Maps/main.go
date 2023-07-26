package main

import "fmt"

func main() {
	fmt.Println("Golang working with maps")
	var temp map[string]string
	temp = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(temp)

	// We can map as implicit like this
	mymap := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(mymap)

	// Declare a map where the values are slices of string
	var m map[string][]string
	fmt.Println(m)
	m = map[string][]string{
		"Coffee": {"Coffee", "Coffe-late"},
		"Tea":    {"Chai", "Latte", "Chai-Latte"},
	}
	fmt.Println(m)
	fmt.Println(m["Coffee"])

	// append
	m["Others"] = []string{"Milk", "Water"} // The ordering is determined by go we have no control
	fmt.Println(m)
	delete(m, "Tea")
	fmt.Println(m)

	// Still we can query it
	v := m["Tea"]
	fmt.Println(v) // This will be an empty array
	// Because it will send default
	// For more control we can do like this
	v, ok := m["Tea"] // ok will be false
	fmt.Println(v, ok)

	// Assignmnet
	c := m
	fmt.Println(c) // Same as m
	c["Others"] = []string{"Hot-milk"}
	fmt.Println(m)
	fmt.Println(c) // Both will be same
}
