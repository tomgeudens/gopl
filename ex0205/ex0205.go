/*
description :
  ex0205 is a variation on the popcount package

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/08/01
*/
package ex0205

func PopCount(x uint64) int {
	var returnValue int = 0

	for ; x > 0; x = x & (x - 1) {
		returnValue++
	}

	return returnValue
}
