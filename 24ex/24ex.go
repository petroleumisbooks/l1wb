package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func DistanceBetweenTwoPoints(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}

func main() {
	a := NewPoint(10, 10)
	b := NewPoint(25, 34)

	fmt.Printf("Distanse between %v and %v is %v\n", a, b, DistanceBetweenTwoPoints(a, b))

}
