package main

import "fmt"

func main() {
	S1 := []int{1, 2, 3, 4, 5} // S1 has length 5 and capacity 5.
	printSlice("S1", S1)
	println("S1 is created by slicing the array [1 2 3 4 5] with a length of 5, so it has 5 elements and a capacity of 5.")

	S2 := []int{6, 7, 8, 9, 10}[:0] // S2 has length 0 and capacity 5.
	printSlice("S2", S2)
	println("S2 is created by slicing the array [6 7 8 9 10] with a length of 0, so it has no elements but a capacity of 5.")
	S3 := S2[:2] // S3 has length 2 and capacity 5. Takes the first 2 elements of S2 (6 7)
	printSlice("S3", S3)
	println("Takes the first 2 elements of S2 (6 7) while keeping the capacity of 5 and the underlying array [6 7 8 9 10].")

	S4 := S3[2:5] // S4 has length 3 and capacity 3. Takes elements from index 2 to 4 (8 9 10)
	printSlice("S4", S4)
	println("Takes elements from index 2 to 4 from S3 (8 9 10)")

	myArray := [3]string{"First", "Second", "Third"}
	fmt.Println("myArray:", myArray)
	mySlice := myArray[:]
	mySlice2 := myArray[:]
	mySlice[0] = "test"
	fmt.Println("myArray:", myArray)
	fmt.Println("First element of mySlice2:", mySlice2[0])
	fmt.Println("First element of myArray:", myArray[0])
	fmt.Println("mySlice and mySlice2 are views of the same array, updating any of them updates the other and the underlying array.")

}

func printSlice(s string, x []int) {
	fmt.Printf("%s length=%d capacity=%d %v\n",
		s, len(x), cap(x), x)
}
