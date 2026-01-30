package main

import "fmt"

// FilterEven возвращает слайс только с чётными числами
func FilterEven(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс и добавляй числа, которые делятся на 2
	var evenNums []int
	for _, v := range nums {
		if v%2 == 0 {
			evenNums = append(evenNums, v)
		}
	}

	return evenNums
}

// FilterPositive возвращает слайс только с положительными числами
func FilterPositive(nums []int) []int {
	// TODO: реализуй функцию
	// Положительные — это числа > 0
	var positiveNums []int
	for _, v := range nums {
		if v > 0 {
			positiveNums = append(positiveNums, v)
		}
	}

	return positiveNums
}

// Double возвращает новый слайс, где каждый элемент умножен на 2
func Double(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс той же длины
	// Заполни его удвоенными значениями

	doubleNums := make([]int, len(nums))
	for i, v := range nums {
		doubleNums[i] = v * 2
	}

	return doubleNums
}

// Reverse возвращает перевёрнутый слайс
func Reverse(nums []int) []int {
	// TODO: реализуй функцию
	// Создай новый слайс и заполни в обратном порядке

	var reverseSlice []int
	for i := len(nums) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, nums[i])
	}

	return reverseSlice
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
