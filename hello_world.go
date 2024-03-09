package main

import (
	"fmt"
	"reflect"
)

func main() {

	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma

	}

	// define slices
	/*
		classics := languages[0:3]
		// modern := make([]string, 4)
		modern := languages[3:7]
		new := languages[7:9]

		fmt.Printf("the classics %v\n", classics)
		fmt.Printf("Basic bois %v\n", modern)
		fmt.Printf("new kids on the block%v\n", new)
	*/

	allLangs := languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[0:4:4]
	frameworks = append(frameworks, "meteor")

	fmt.Printf("all frameworks %v\n", frameworks)
	fmt.Printf("js frameworks %v\n", jsFrameworks)
}
