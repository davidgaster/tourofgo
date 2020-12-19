package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(Sqrt(3))

}

func Sqrt(x float64) float64 {
	var z float64
	if x >= 4  {
		z = x/2
	} else {
		z = x
	}
	for x < z*z {
		z -= (z*z - x) / (2*z)
	}
	return z
}