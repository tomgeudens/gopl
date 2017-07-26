/*
description :
  ftoc does a couple of temperature conversions using functions

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/26
*/
package main

import "fmt"

func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func ftoK(f float64) float64 {
	return ctoK(ftoC(f))
}

func ctoF(c float64) float64 {
	return (c * 9 / 5) + 32
}
func ctoK(c float64) float64 {
	return c + 273.15
}

func ktoC(k float64) float64 {
	return k - 273.15
}
func ktoF(k float64) float64 {
	return ctoF(ktoC(k))
}

func main() {
	const freezingF = 32.0
	const boilingF = 212.0
	const freezingC = 0.0
	const boilingC = 100.0
	const freezingK = 273.15
	const boilingK = 373.15

	fmt.Printf("freezing : %g°F = %g°C = %gK\n", freezingF, ftoC(freezingF), ftoK(freezingF))
	fmt.Printf("boiling  : %g°F = %g°C = %gK\n", boilingF, ftoC(boilingF), ftoK(boilingF))

	fmt.Printf("freezing : %g°C = %g°F = %gK\n", freezingC, ctoF(freezingC), ctoK(freezingC))
	fmt.Printf("boiling  : %g°C = %g°F = %gK\n", boilingC, ctoF(boilingC), ctoK(boilingC))

	fmt.Printf("freezing : %gK = %g°C = %g°F\n", freezingK, ktoC(freezingK), ktoF(freezingK))
	fmt.Printf("boiling  : %gK = %g°C = %g°F\n", boilingK, ktoC(boilingK), ktoF(boilingK))
}
