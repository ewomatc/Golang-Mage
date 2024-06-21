package main

import (
	"fmt"
	"strings"
)

func main() {
	// string operations

	// length of a string
	randomText := "This is my random text"
	fmt.Println(len(randomText))

	// string comparion - case insensitive
	string1 := "Mage's lullaby"
	string2 := "mage's lullaby"
	fmt.Println(strings.EqualFold(string1, string2))

	// print the index of a character in a string
	strindex := "Print, My 👀 index"
	fmt.Println(strings.Index(strindex, "👀"))

	// create substrings from a string by slicing that string
	// pass in the start indes and end index to the string you want to slice
	startIndex := strings.Index(strindex, ",")
	substring := strindex[startIndex:]
	fmt.Println(substring)

	// replace a word or character occurence ina string, note: these only apply to the first occurence
	stringToMutate := "Hello World"
	fmt.Println("original string", stringToMutate)
	mutatedString := strings.Replace(stringToMutate, "World", "Mage", 1) //replace() takes 4 args, pass in the name of the string to mutate, then the current word, then the word to replace it with, and finally the index of the current word, or character
	fmt.Println("Mutated string", mutatedString)

	// change string casing
	lowerstr := "i am mage"
	upperStr := strings.ToUpper(lowerstr)
	fmt.Println(upperStr) // works both ways

	// check is a string contains a character or word
	fullStr := "Go Programming"
	fmt.Println(strings.Contains(fullStr, "Go"))
	fmt.Println(strings.Contains(fullStr, "Golang"))
}
