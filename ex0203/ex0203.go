/*
description :
  ex0203 is a variation on the popcount package

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/31
*/
package ex0203

// pc covers 8 bits and pc[i] contains the population count of i
var pc [256]byte

func init() {
	for i, _ := range pc { // only need the index
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var returnValue int = 0

	for i := 0; i < 8; i++ {
		returnValue += int(pc[byte(x>>(uint(i)*8))])
	}

	return returnValue
}
