package polynomial

import (
	"fmt"
	"testing"
)

func TestPolynomial_Print(t *testing.T) {
	polyn := &polynomial{coefficients: []uint8{1, 2, 3}}
	if polyn.String() != "1x^2 + 2x + 3" {
		t.Fatal()
	}
	fmt.Println(polyn.String())
}
