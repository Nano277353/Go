package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Nano"
	var length = len(name) //4
	fmt.Println(name[0:2]) //te
	fmt.Println(length)    //4
	var name2 = name[:]
	var apellido = "Ramirez"
	fmt.Println(name2) //Nano Copy
	// Strings are immutable
	fmt.Println(strings.Join([]string{name, apellido}, " "))
}
