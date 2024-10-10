package main

import "fmt"

// Go has various types including strings, intergers, floats,booleans, etc

// Strings can be added together with +
// Integers and floats
// Booleans, with boolean operators as you would expect

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
