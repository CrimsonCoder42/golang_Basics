//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	Base  int
	Side1 int
	Side2 int
}

func triArea(a, c, b int) float64 {
	totals := (a + b + c) * (-a + b + c) * (a - b + c) * (a + b - c)
	area := 0.25 * math.Sqrt(float64(totals))
	return area
}

func triPerim(s1, s2, base float32) float32 {
	perimeter := s1 + s2 + base
	return perimeter
}

func main() {
	triangle1 := Triangle{3, 3, 3}
	triangle1Result := triArea(triangle1.Base, triangle1.Side1, triangle1.Side2)

	fmt.Println(math.Round(triangle1Result))
}
