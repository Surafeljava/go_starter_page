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
	perimeter() float64
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
	for _, s := range shp {
		total += s.area()
	}
	fmt.Println("Area: ", total)
	return total
}

func totalPerimeter(shp ...shape) float64 {
	perim := 0.0
	for _, s := range shp {
		perim += s.perimeter()
	}
	fmt.Println("Perimeter: ", perim)
	return perim
}

func other() {
	c1 := new(circle)
	c1.radius = 2.5

	r1 := new(rectangle)
	r1.width = 10
	r1.height = 15

	totalArea(c1)
	totalPerimeter(c1)

	totalArea(r1)
	totalPerimeter(r1)

}
