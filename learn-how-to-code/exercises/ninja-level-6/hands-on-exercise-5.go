package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return 2 * c.radius * math.Pi
}

func main() {
	s1 := square{
		length: 1.97,
	}

	c1 := circle{
		radius: 1.97,
	}

	fmt.Printf("Area of square with length %v is %v", s1.length, s1.area())
	fmt.Printf("\nArea of circle with radius %v is %v", c1.radius, c1.area())
}
