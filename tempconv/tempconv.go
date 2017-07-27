/*
description :
  tempconv is a package that provides temperature conversion functions

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/27
*/
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const AbsoluteZeroC Celsius = -273.15
const AbsoluteZeroF Fahrenheit = -459.67
const AbsoluteZeroK Kelvin = 0.0
const FreezingC Celsius = 0.0
const FreezingF Fahrenheit = 32.0
const FreezingK Kelvin = 273.15
const BoilingC Celsius = 100.0
const BoilingF Fahrenheit = 212.0
const BoilingK Kelvin = 373.15

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
