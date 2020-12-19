package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var z float64
	if x >= 4  {
		z = x/2
	} else {
		z = x
	}
	for x < z*z {
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}

type ErrNegativeSqrt float64
func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(err))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
