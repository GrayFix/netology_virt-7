package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a length in m: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Printf("Length in ft: %.4f\n", m_to_ft(input))
}

func m_to_ft(m float64) float64 {
	return m / 0.3048
}
