package main

import (
	"fmt"
	"github.com/williamkoller/introduction/adresses"
)

func main() {
	typeAdresses := adresses.TypeAdresses("Rodovia dos Imigrantes")

	fmt.Println(typeAdresses)
}
