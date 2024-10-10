// In go variables are explicitly declared and used by the
// compiler to e.g. check type-correctedness of function calls

// var declares 1 or more variables
// you can declare multiple variables at once
// Go will infer the type of inialized variables

// Variables declared without a corresponding initialization are zero-valued.
// For example the 0 value for an int is 0

// The := syntax is shorthand for declaring and initializing a variable
// e.g. for var f string = "apple" in this case, the syntax is only availiable
// inside functions

package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
