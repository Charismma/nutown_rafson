package main

import (
	"fmt"
	"math"
)

func proizvodnaya(x float64) float64 {
	return (((3 * x * x) / 5) - (2 * x) - 1)
}
func sama_function(x float64) float64 {
	return (((x * x * x) / 5) - (x * x) + x)
}
func findZero(initial_guess float64, maxError float64) float64 {
	var x float64 = initial_guess
	for i := 1; i <= 100; i++ {
		var y float64 = sama_function(x)
		if math.Abs(y) < maxError {
			break
		}
		x = (x - y) / proizvodnaya(x)
	}
	return x
}

func main() {
	var start float64 = -10
	maxError := 0.1
	for i := start; i <= 10; i += 0.1 {
		b := findZero(i, maxError)
		fmt.Println("Корень при начальном значении: ", i, " равен: ", b)
		fmt.Println("")
		fmt.Println("")
	}
}
