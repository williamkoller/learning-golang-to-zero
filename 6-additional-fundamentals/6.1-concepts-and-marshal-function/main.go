package main

import (
	"bytes"
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
	newDog := dog{"aderbado", "linguica", 1}

	fmt.Println(newDog)

	dogJSON, err := json.Marshal(newDog)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dogJSON)

	bytes.NewBuffer(dogJSON)

	fmt.Println(bytes.NewBuffer(dogJSON))

	dogTwo := map[string]string{
		"name": "Toby",
		"breed": "Poodle",
	}

	dogTwoJSON, err := json.Marshal(dogTwo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dogTwoJSON)
	fmt.Println(bytes.NewBuffer(dogTwoJSON))

}
