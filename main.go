package main

import "fmt"

func main() {

	colors := map[string]string {
		"red" : "#ff0000",
		"green" : "#ff2343",
		"white" : "000000",
		"black" : "ffffff",
	}

	printMap(colors)
}

func printMap (c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for", color, "is", hex)
	}
}