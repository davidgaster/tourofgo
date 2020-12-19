package main

import "fmt"

// Array's length CANNOT be changed as it is part of its type.
//
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var slice []int = primes[1:4]
	fmt.Println(slice)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(names)
	fmt.Println(a, b)
}
