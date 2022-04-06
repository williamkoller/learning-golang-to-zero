package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("saving user data %s in database\n", u.name)

}

func (u user) major() bool {
	return u.age > 18
}

func (u *user) haveABirthday() {
	u.age++
}

func main() {
	userOne := user{"William", 32}
	userOne.save()

	userTwo := user{"Barbara", 17}
	userTwo.save()

	major := userTwo.major()
	fmt.Println(major)

	userTwo.haveABirthday()
	fmt.Println(userTwo.age)
}