package main

import "fmt"

// FilterEven возвращает слайс только с чётными числами
func FilterEven(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс и добавляй числа, которые делятся на 2
	return nil
}

// FilterPositive возвращает слайс только с положительными числами
func FilterPositive(nums []int) []int {
	// TODO: реализуй функцию
	// Положительные — это числа > 0
	return nil
}

// Double возвращает новый слайс, где каждый элемент умножен на 2
func Double(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс той же длины
	// Заполни его удвоенными значениями
	return nil
}

// Reverse возвращает перевёрнутый слайс
func Reverse(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс и заполни в обратном порядке
	return nil
}

func main() {
	nums := []int{-3, -1, 0, 1, 2, 3, 4, 5}

	fmt.Println("=== Фильтрация и преобразование ===")
	fmt.Println("Исходный слайс:", nums)

	fmt.Println("\nЧётные числа:", FilterEven(nums))
	fmt.Println("Положительные:", FilterPositive(nums))
	fmt.Println("Удвоенные:", Double(nums))
	fmt.Println("Перевёрнутый:", Reverse(nums))

	fmt.Println("\nОригинал не изменился:", nums)
}
