// Go supports constants of character, string, boolean and numneric values

// const declares a constant value

// A const statement can appear anywhere a var statement can

// Constant expressions perform arithmetic with arbitrary precision

// A numeric constant has no type until it's given one, such as by an
// explicit type conversion

// A number can be given a type by using it in a context that requires one,
// such as variable assignment or function call.
// For example, here math.Sin expects a float64

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500_000_000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
