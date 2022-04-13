package adresses

import "strings"

// TypeAdresses verify if a address have a type valid and return
func TypeAdresses(address string) string {
	typesValids := []string{"rua", "avenida", "estrada", "rodovia"}

	lowercaseAddress := strings.ToLower(address)

	firstWordOfAddress := strings.Split(lowercaseAddress, " ")[0]

	addressHaveTypeValid := false

	for _, typesValid := range typesValids {
		if typesValid == firstWordOfAddress {
			addressHaveTypeValid = true
		}
	}

	if addressHaveTypeValid {
		return strings.Title(firstWordOfAddress)
	}

	return "type invalid"
}
