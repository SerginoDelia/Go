// Maps are built-in associative data type (sometimes called
// hashes, or dicts in other languages).

package main

import (
	"fmt"
	"maps"
)

func main() {
	// to create an empty map use the built-in make
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all its
	// key/value pairs
	fmt.Println("map:", m)

	// Get the value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, the 0 value of the value type is returned
	// v3 := m["k3"]
	// fmt.Println("v3:", k3)

	// The built-in len returns the number of key/value
	// pairs when called on a map
	fmt.Println("len:", len(m))

	// The built-in delete removes key/value pairs from a map
	delete(m, "k2")

	// To remove all key/value pairs from a map, use the clear buit-in
	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map. This can be used to
	// disambiguate between missing keys and keys with zero values like 0
	// or "". Here we didn't need the value itself, so we ignored it with
	// the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line
	// using this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// the maps package contains a number of useful utility functions for maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
