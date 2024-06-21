package main

import "fmt"

func main() {
	// strings, ':=' type inference
	name := "Mage"
	fmt.Println("Hello", name)

	var agency string = "Pipdevs"
	fmt.Println("Welcome to", agency)

	// integers,
	// '%v is placeholder
	// \n newline since print() doesnt automatically add newline
	totalCars := 50
	fmt.Printf("We have %v cars in our fleet \n", totalCars)

	// boolean
	insuranceIncluded := true
	fmt.Println("Is Insurance Included?", insuranceIncluded)

}
