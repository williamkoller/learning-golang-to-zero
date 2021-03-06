package adresses_test

import (
	"testing"

	"github.com/williamkoller/introduction/adresses"
)

type testAddress struct {
	addressInsert  string
	returnExpected string
}

func TestTypeAdresses(t *testing.T) {
	t.Parallel()
	cenaryTest := []testAddress{
		{
			"Rua ABC", "Rua",
		},
		{
			"Avenida Paulista", "Avenida",
		},
		{
			"Rodovia Imigrantes", "Rodovia",
		},
		{
			"Rodovia do Maraja", "Rodovia",
		},
		{
			"Estrada da Graciosa", "Estrada",
		},
		{
			"Estrada dos Bobos", "Estrada",
		},
		{
			"", "type invalid",
		},
	}

	for _, cenaryT := range cenaryTest {
		typeOfAddressReceived := adresses.TypeAdresses(cenaryT.addressInsert)
		if typeOfAddressReceived != cenaryT.returnExpected {
			t.Errorf("type received is different to expected! expected %s and received %s",
				typeOfAddressReceived,
				cenaryT.returnExpected)
		}
	}
}

func TestAny(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Test Fail")
	}
}
