package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world.")
	fmt.Println(strings.Repeat("-", 50))

	// variable and constants
	var userName string = "Jeter"
	fmt.Println(userName)
	fmt.Println(strings.Repeat("-", 50))

	const number int = 123
	fmt.Println(number)
	fmt.Println(strings.Repeat("-", 50))

	// formatted
	var userName2 = "Ray"
	const years = 10
	fmt.Println("name:", userName2, "years old:", years)
	fmt.Printf("hello my name is %v, I'm %v years old.\n", userName2, years)
	fmt.Println(strings.Repeat("-", 50))

	// data type
	var userName3 string
	var number2 int
	userName3 = "hello world."
	number2 = 123
	fmt.Printf("name: %v number: %v\n", userName3, number2)
	fmt.Println(strings.Repeat("-", 50))
}
