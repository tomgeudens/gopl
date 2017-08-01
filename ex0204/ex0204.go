/*
description :
  ex0204 is a variation on the popcount package

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/08/01
*/
package ex0204

func PopCount(x uint64) int {
	var returnValue int = 0

	for x > 0 {
		if x&1 == 1 { // this tests if a number is odd
			returnValue++
		}
		x >>= 1 // bit shift to the right
	}

	return returnValue
}
