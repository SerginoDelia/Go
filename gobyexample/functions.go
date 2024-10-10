// Functions are central in Go.

package main

import "fmt"

// Here a function that takes 2 ints and return their sum as an int
func plus(a int, b int) int {

	// Go require explicit returns, i.e. it won't automatically return
	// the value of the last expression
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to
// the final parameter that declares the type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Call a function just as you'd expect, with the name(args)
func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	fmt.Println(plus(5, 8))
}
