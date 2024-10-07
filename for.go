// for is Go's only looping construct.
// Here are some basic examples of for loops

// The most basic type, with a single condition

// A classic initial/condition/after for loop

// Another way of accoplishing the basic "do this N times"
// iteration is range over an integer

// for without a condition will loop repeatedly until you break out
// of the loop or return from the enclosing function

// You can also continue to the next iteraion of the loop

package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
