/*
description :
  cfk converts its numeric arguments to Celsius, Fahrenheit and Kelvin

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/27
*/
package main

import "fmt"
import "os"
import "strconv"

import "github.com/tomgeudens/gopl/tempconv"

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			continue
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s = %s\n", c, tempconv.CtoF(c), tempconv.CtoK(c))
		fmt.Printf("%s = %s = %s\n", f, tempconv.FtoC(f), tempconv.FtoK(f))
		fmt.Printf("%s = %s = %s\n", k, tempconv.KtoC(k), tempconv.KtoF(k))

		fmt.Println()
	}
}
