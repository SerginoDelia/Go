// Variadic functions can be called with any number of
// trialing arguments. For example, fmt.Println is common
// variadic function

package main

import "fmt"

// Here's a function that will take an arbitrary number
// of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// Within the function, the type of nums is equvalent to
	// []int. We can call len(nums), to iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// variadic functions can be called in the usual way with
	// individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// If you have multiple args in a slice, apply them to a variadic
	// function using the fun(slice...) like this

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
