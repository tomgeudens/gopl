/*
description :
  boiling prints the boiling point of water

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/26
*/
package main

import "fmt"

const boilingF = 212.0
const boilingC = 100.0
const boilingK = 373.15

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	var k = (f-32)*5/9 + 273.15

	fmt.Printf("boiling point = %g°F or %g°C or %gK\n", f, c, k)

	c = boilingC
	f = (c * 9 / 5) + 32
	k = c + 273.15

	fmt.Printf("boiling point = %g°C or %g°F or %gK\n", c, f, k)

	k = boilingK
	c = k - 273.15
	f = (k-273.15)*9/5 + 32

	fmt.Printf("boiling point = %gK or %g°C or %g°F\n", k, c, f)
}
