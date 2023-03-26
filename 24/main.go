package main

/*Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором.*/

import (
	"fmt"
	"math"
)

// Создаем именованный тип Point
type Point struct {
	//инициализируем поля
	x float64
	y float64
}

// Создаем метод-конструктор для Point
func (point *Point) Init(x, y float64) {
	point.x = x
	point.y = y
}

// Создаем метод и измеряем расстояние по формуле
func (pointA *Point) measure(pointB *Point) float64 {
	return math.Sqrt((math.Pow((pointB.x-pointA.x), 2) + math.Pow((pointB.y-pointA.y), 2)))
}
func main() {
	//Объявляем и инициализируем первую точку
	var pointA Point
	pointA.Init(3, 2)
	//Объявляем и инициализируем вторую точку
	var pointB Point
	pointB.Init(9, 7)
	fmt.Printf("PointA = %+v\n", pointA)
	fmt.Printf("PointB = %+v\n", pointB)
	//Измеряем расстояние между точкой А и B
	fmt.Printf("Distance = %v", pointA.measure(&pointB))
}
