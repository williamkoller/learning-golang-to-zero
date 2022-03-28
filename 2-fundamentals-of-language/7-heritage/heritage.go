package main

import "fmt"

// this is the closest to inheritance in go

type people struct {
	name string
	lastName string
	age int8
	height int16
}

type student struct {
	people
	course string
	college string
}

func main() {
	p := people{"William", "Koller", 32, 182}
	fmt.Println(p)

	e := student{p, "ADS", "Faculdade Integradas Camoes"}
	fmt.Println(e)
}