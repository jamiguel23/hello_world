package main

import "fmt"

func main() {

	//defining an array that stores 4 items
	var testArr = [4]string{"yes", "no", "stop", "naur"}

	// loop through
	fmt.Println("for loop")
	for i := 0; i < len(testArr); i++ {
		item := testArr[i]
		fmt.Println(i, item)
	}

	var testArr1 = [...]string{"what", "life", "donut", "test"}

	fmt.Println("range")

	for index, item := range testArr1 {
		fmt.Println(index, item)
	}
}
