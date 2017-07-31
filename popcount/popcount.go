/*
description :
  popcount is a package that provides a function to determine the population
  count (number of set bits) of an integer number

author :
  Tom Geudens (https://github.com/tomgeudens/)

modified :
  2017/07/31
*/
package popcount

// pc covers 8 bits and pc[i] contains the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return (int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]))

}
