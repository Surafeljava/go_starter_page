package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (r rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func totalArea(shp ...shape) float64 {
	total := 0.0
	for s := range shp {
		total += shp[s].area()
	}
	fmt.Println(total)
	return total
}

func main() {
	// c1 := new(circle)
	// c1.radius = 2.5
	// fmt.Println(c1.area())
	// totalArea(c1)

	fmt.Println("Hello")
}
