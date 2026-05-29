package main

import "fmt"

func main() {
	var age int
	var name string
	var drinking bool

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)
	if age > 18 {
		drinking = true
	}

	fmt.Println("Name:", name, "Age:", age, "Drinking:", drinking)
}
