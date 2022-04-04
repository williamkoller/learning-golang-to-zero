package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("increment i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("increment j", j)
		time.Sleep(time.Second)

	}

	names := [3]string{"William", "Barbara", "Pietro"}

	for indice, name := range names {
		fmt.Println(indice, name)
	}


	for indice, letra := range "DEVELOPER" {
		// print table ascii
		fmt.Println(indice, letra)

		fmt.Println(indice, string(letra))
	}

	user := map[string]string{
		"name": 		"William",
		"lastName": "Koller",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}