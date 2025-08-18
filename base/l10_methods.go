package main

import (
	"fmt"
	"math"
)

type Complex struct {
	X, Y float64
}

func (c Complex) Abs() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

// pointer receivers
func (c *Complex) Scale(f float64) {
	c.X = c.X * f
	c.Y = c.Y * f
}

// pointer and functions
func Scale2(c *Complex, s float64) {
	c.X = c.X * s
	c.Y = c.Y * s
}

func main() {
	v := Complex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
	Scale2(&v, 10)
	fmt.Println(v.Abs())

	p := &Complex{3, 4}
	p.Scale(3)
	Scale2(p, 10)
	fmt.Println(v, p)
}
