package means_test

import (
	"math"
	"testing"

	. "github.com/williamkoller/test-advanced/means"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		areaExpected := float64(120)
		areaReceived := rect.Area()

		if areaExpected != areaReceived {
			t.Fatalf("The Area received %f is different of expected %f",
				areaReceived,
				areaExpected,
			)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		cir := Circle{10}

		areaExpected := float64(math.Pi * 100)
		areaReceived := cir.Area()

		if areaExpected != areaReceived {
			t.Fatalf("The Area received %f is different of expected %f",
				areaReceived,
				areaExpected,
			)
		}
	})
}
