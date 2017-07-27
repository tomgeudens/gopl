/*
description :
  tempconv is a package that provides temperature conversion functions

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/27
*/
package tempconv

// CtoF converts a Celsius temperature to Fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

// CtoK converts a Celsius temperature to Kelvin
func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FtoC converts a Fahrenheit temperature to Celsius
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FtoK converts a Fahrenheit temperature to Kelvin
func FtoK(f Fahrenheit) Kelvin {
	return Kelvin(CtoK(FtoC(f)))
}

// KtoC converts a Kelvin temperature to Celsius
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KtoF converts a Kelvin temperature to Fahrenheit
func KtoF(k Kelvin) Fahrenheit {
	return Fahrenheit(CtoF(KtoC(k)))
}
