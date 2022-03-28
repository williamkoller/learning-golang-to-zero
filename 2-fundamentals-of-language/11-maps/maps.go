package main

import "fmt"

func main() {
	user := map[string]string{
		"name":    "William",
		"surname": "Koller",
	}

	fmt.Println(user["name"])

	userTwo := map[string]map[string]string{
		"name": {
			"first": "Will",
			"last":  "Koller",
		},
	}

	fmt.Println(userTwo)

	delete(userTwo["name"], "first")

	fmt.Println(userTwo)

	userTwo["signo"] = map[string]string{
		"name": "cancer",
	}

	fmt.Println(userTwo)

}
