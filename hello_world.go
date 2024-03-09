package main

import "fmt"

func main() {

	// user specified types
	const a int32 = 12        // 32 bit
	const b float32 = 34.2    // 32 bit float
	var c complex128 = 1 + 4i // 128 complex number
	var d uint16 = 14         // 16 bit unsigned integer

	// default types
	n := 42                    // int
	pi := 3.14                 // float54
	x, y := true, false        // bool
	z := "this is a Go string" // string

	fmt.Printf("user-specified \n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default \n %T %T %T %T %T\n", n, pi, x, y, z)
}
