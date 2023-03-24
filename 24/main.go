package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func (point *Point) Init(x, y float64) {
	point.x = x
	point.y = y
}
func (pointA *Point) measure(pointB *Point) float64 {
	return math.Sqrt((math.Pow((pointB.x-pointA.x), 2) + math.Pow((pointB.y-pointA.y), 2)))
}
func main() {
	var pointA Point
	pointA.Init(3, 2)
	var pointB Point
	pointB.Init(9, 7)
	fmt.Printf("PointA = %+v\n", pointA)
	fmt.Printf("PointB = %+v\n", pointB)
	fmt.Printf("Distance = %v", pointA.measure(&pointB))
}
