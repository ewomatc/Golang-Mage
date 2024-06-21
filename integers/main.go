package main

import (
	"fmt"
)

func main() {

	// convert 288.15 kelvin to farenheit
	tempKelvin := 288.15
	tempFaren := (tempKelvin-273.15)*1.8 + 32
	fmt.Printf("%v kelvin in farenheit is %v \n", tempKelvin, tempFaren)

	// convert 70 farenheit to kelvin
	tempFaren2 := 70.0
	tempKelvin2 := (tempFaren2-32)/1.8 + 273.15
	fmt.Printf("%v farenheit in kelvin is %v \n", tempFaren2, tempKelvin2)

}
