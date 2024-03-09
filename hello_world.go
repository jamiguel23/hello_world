package main

import (
	"fmt"
)

func average(x []float64) (avg float64) {

	// initialize floating variable to 0
	total := 0.0
	// check if length is 0
	if len(x) == 0 {
		avg = 0
	} else {
		// for loop to iterate through the range of the array
		for _, v := range x {
			// add float to total
			total += v
		}
		// divide total with lengh of the array but as a float64 (the same data type)
		avg = total / float64(len(x))
	}

	return
}
func main() {

	// x := []float64{2.32, 23.42, 42.3, 29, 5}
	z := []float64{}

	fmt.Println(average((z)))
}
