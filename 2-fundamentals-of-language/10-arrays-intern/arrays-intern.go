package main

import "fmt"

func main() {
	// arrays interns

	slice := make([]float32, 10, 11)
	fmt.Println(slice)

	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Println(slice)
	
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	sliceTwo := make([]float32, 5)
	fmt.Println(sliceTwo)

	sliceTwo = append(sliceTwo, 10)

	fmt.Println(len(sliceTwo))
	fmt.Println(cap(sliceTwo))

}
