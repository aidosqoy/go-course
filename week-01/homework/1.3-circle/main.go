package main

import (
	"fmt"
	"math"
)

// Circle вычисляет площадь и периметр круга по радиусу
func Circle(radius float64) (area, perimeter float64) {
	// TODO: реализуй функцию
	// Используй math.Pi для числа π

	area = math.Pi * (radius * radius)
	perimeter = 2 * math.Pi * radius
	return area, perimeter
}

func main() {
	a, p := Circle(5)
	fmt.Printf("Радиус: 5\n")
	fmt.Printf("Площадь: %.2f\n", a)
	fmt.Printf("Периметр: %.2f\n", p)

	// Проверь на других значениях
	a2, p2 := Circle(10)
	fmt.Printf("\nРадиус: 10\n")
	fmt.Printf("Площадь: %.2f\n", a2)
	fmt.Printf("Периметр: %.2f\n", p2)

	a3, p3 := Circle(15)
	fmt.Printf("\nРадиус: 15\n")
	fmt.Printf("Площадь: %.2f\n", a3)
	fmt.Printf("Периметр: %.2f\n", p3)
}
