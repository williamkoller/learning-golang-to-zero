package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dogJSON := `{"name":"aderbado","breed":"linguica","age":1}`

	var d dog

	if err := json.Unmarshal([]byte(dogJSON), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)

	dogTwoJSON := `{"name":"aderbado","breed":"linguica"}`

	dTwo := make(map[string]string)

	if err := json.Unmarshal([]byte(dogTwoJSON), &dTwo); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dTwo)
}
