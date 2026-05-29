package main

import "fmt"

func main() {
	var myArray = [3]string{"Houseki", "no", "Kuni"}
	myArrayCopy := myArray
	fmt.Print(myArray[1]) //"no"
	myArray[1] = "of"
	fmt.Print(myArrayCopy[1]) //Keeps "no"
	fmt.Println(myArray)      //[Houseki of Kuni]
	fmt.Println(myArrayCopy)  //[Houseki no Kuni]
}
