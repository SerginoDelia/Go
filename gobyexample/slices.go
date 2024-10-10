// Slices are more important data in Go, giving a more
// powerful interface to sequence than arrays

package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays slices are typed only by the elements they contain
	// (not the number of elements). An uninitilized slice equals to nil
	// and has a length of 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use the builtin make
	// Here we make a slice of strings of length 3 (initially zero-valued).
	// By default a new slice's capacity is equal to its length; if we know
	// the slice is going to grow ahead of time, it's possible to pass a capacity
	// explicitly as an additional parameter to make
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// we can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several
	// more that make them richer than arrays. One is the builtin append,
	// which returns a slice containing one or more values. Note that we
	// we need to accept a return value from the append as we may get a
	// new slice value

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can be copy'd. Here we create an empty slice c of the same
	// length as s and copy into from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high].
	// For example, this get a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slices up to but excluding s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slices up from (and including) s[2]
	l = s[2:]
	fmt.Print("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices package contains a number of useful utility functions for slices
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can be composed of multi-demensional data structures. The length of an
	// inner slices can very, unlike multi-demensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := 1 + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
