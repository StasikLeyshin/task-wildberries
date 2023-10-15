/*

Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64 // Координата x
	y float64 // Координата y
}

// NewPoint - конструктор для структуры Point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// DistanceTwoPoints - определение расстояния между двумя точками
func (point1 *Point) DistanceTwoPoints(point2 *Point) float64 {
	dx := point1.x - point2.x
	dy := point1.y - point2.y

	return math.Sqrt(dx*dx + dy*dy)
}

func main() {

	// Исходные точки
	point1 := NewPoint(0.0, 8.0)
	point2 := NewPoint(8.0, 0.0)

	fmt.Println("Distance:", point1.DistanceTwoPoints(point2))
}
