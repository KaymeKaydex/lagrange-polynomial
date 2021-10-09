package polynomial

import (
	"fmt"
	"strings"
)

type polynomial struct {
	coefficients []uint8
}

func New(coefficients []uint8) *polynomial {
	return &polynomial{coefficients: coefficients}
}

func (p *polynomial) String() string {
	coefficientsCount := len(p.coefficients) - 1
	polynomialArr := make([]string, 0, coefficientsCount)
	for index, coef := range p.coefficients {
		switch coefficientsCount - index {
		case 0:
			polynomialArr = append(polynomialArr, fmt.Sprintf("%d", coef))
		case 1:
			polynomialArr = append(polynomialArr, fmt.Sprintf("%dx", coef))
		default:
			polynomialArr = append(polynomialArr, fmt.Sprintf("%dx^%d", coef, coefficientsCount-index))
		}
	}

	return strings.Join(polynomialArr, " + ")
}
