package main

import "fmt"

/**
struct is collection fields
*/

type user struct {
	name string
	age  int8
	addressUser address
}

type address struct {
	publicPlace string
	number uint8
}

func main() {
	var u user

	u.name = "William"
	u.age = 32
	fmt.Println(u)

	addressA := address{"Rua Luziano Cord", 120}

	userTwo := user{"Koller", 33, addressA}
	fmt.Println(userTwo)

	userThree := user{name: "Barbara"}
	fmt.Println(userThree)

	userFour := user{age: 34}
	fmt.Println(userFour)
}
