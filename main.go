package main

import "fmt"

func main() {
	// First way of defining map
	colors := map[string]string {
		"red" : "#ff0000",
		"green" : "#ff2343",
	}

	// Second way of defining map
	var vehicles map[string]int

	// Third way of defining map
	electronics := make(map[string]string)

	electronics["mobile"] = "android"
	electronics["laptop"] = "mac"


	fmt.Println(colors)
	fmt.Println(vehicles)
	fmt.Println(electronics)

	delete(electronics, "mobile")
	fmt.Println(electronics)
}