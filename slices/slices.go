package main

import "fmt"

func main() {
	S1 := make([]int, 5) // S1 has length 5 and capacity 5. Default values are 0 and are all shown.
	printSlice("S1", S1)

	S2 := make([]int, 0, 5) // S2 has length 0 and capacity 5. Default values are 0 but are not shown.
	printSlice("S2", S2)

	S3 := S2[:2] // S3 has length 2 and capacity 5. Takes the first 2 elements of S2 (0 0)
	printSlice("S3", S3)

	S4 := S3[2:5] // S4 has length 3 and capacity 3. Takes elements from index 2 to 4 (0 0 0)
	printSlice("S4", S4)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
