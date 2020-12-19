package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum, x1, x2 := 0, 0, 1
	
	return func() int {
		num := sum
		sum = x1 + x2
		x1 = x2
		x2 = sum
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
