package main

import (
	"fmt"

	"github.com/williamkoller/2-package/auxiliary"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing file main")
	auxiliary.Write()

	err := checkmail.ValidateFormat("william@mail.com")
	fmt.Println(err)
}
