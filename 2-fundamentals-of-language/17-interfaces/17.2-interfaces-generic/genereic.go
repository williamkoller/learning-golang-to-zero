package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("William")
	generic(32)
	generic(true)
	generic(99.98)

	myMap := map[interface{}]interface{} {
		1: "String",
		float32(100): true,
		"String": "String",
		true: float32(1002),
	}

	fmt.Println(myMap)
}