/*
Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v2 := Vertex2{3, 4}
	fmt.Println(Abs2(v2))
}

/*
Methods are functions
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
*/
type Vertex2 struct {
	X, Y float64
}

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
