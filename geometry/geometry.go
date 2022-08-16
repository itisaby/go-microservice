package geometry

import "math"

//Area
func Area(x, y float64) float64 {
	area := x * y
	return area
}

// Diagonal Function
func Diagonal(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
