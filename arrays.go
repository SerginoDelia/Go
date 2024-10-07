// In Go, an array is a numbered sequence of elements of a
// specific length. In typical Go code, slices are much more
// common; arrays are useful in special scenarios

package main

import "fmt"

func main() {
	// Here we create an array that will hold exactly 5 ints. The
	// type of elements and length are both part of the arrays's type.
	// By default an array is a zero-valued, which for ints means 0s

	var a [5]int
	fmt.Println("emp:", a)

	// We can set the value at an index using the array[index] = value
	// syntax, and get the value with array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// the built-in len returns the length of an array
	fmt.Println("len:", len(a))

	// use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// you can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the element in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Array types are one-demensional, but you can compose types to build
	// multidemesional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// You can create and initialize multi-demensional arrays at once too.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
