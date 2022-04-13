package adresses_test

import (
	"testing"

	"github.com/williamkoller/introduction/adresses"
)

func TestTypeAdresses(t *testing.T) {
	adressesTest := "Avenida Paulista"

	typeAdressExpected := "avenida"

	typeAdresseReceived := adresses.TypeAdresses(adressesTest)

	if typeAdresseReceived != typeAdressExpected {
		t.Errorf("type received is different to expected! expected %s and received %s",
			typeAdressExpected,
			typeAdresseReceived)
	}
}
