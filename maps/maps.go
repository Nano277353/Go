package main

import "fmt"

func main() {
	fmt.Println("Maps in Go are the equivalent of dictionaries in Python or objects in JavaScript.")
	agesMap := make(map[string]int)
	agesMap["Alice"] = 21
	agesMap["Nano"] = 20
	agesMap["Bernie"] = 16
	agesMap["Leo"] = 20
	agesMap["Regina"] = 20
	fmt.Println("Ages Map:", agesMap)

	age := agesMap["Nano"]
	fmt.Println("Nano's age:", age)
	delete(agesMap, "Alice")
	fmt.Println("Ages Map after deleting Alice:", agesMap)

	experienceMap := map[string]int{"Kris": 39, "Susie": 25, "Asriel": 20}
	fmt.Println("Experience Map:", experienceMap)
}
