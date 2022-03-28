package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arrayOne [5]int

	arrayOne[0] = 12
	arrayOne[1] = 2
	arrayOne[2] = 22
	arrayOne[3] = 26
	arrayOne[4] = 29
	fmt.Println(arrayOne)

	arrayTwo := [5]string{"text one", "text two", "text three", "text four", "text five"}
	fmt.Println(arrayTwo)

	arrayThree := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayThree)

	slice := []int{10, 11, 12, 13, 14, 22, 33, 44, 77}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arrayThree))

	slice = append(slice, 18)
	fmt.Println(slice)

	sliceTwo := arrayTwo[1:3]
	fmt.Println(sliceTwo)

	arrayTwo[1] = "altered position"
	fmt.Println(sliceTwo)
}
