package main

import "fmt"

// IsPrime проверяет, является ли число простым
func IsPrime(n int) bool {
	// TODO: реализуй функцию
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Тесты
	testCases := []int{-5, 0, 1, 2, 3, 4, 5, 7, 11, 13, 15, 17, 18, 19, 20, 23, 100}

	for _, n := range testCases {
		result := IsPrime(n)
		fmt.Printf("IsPrime(%d) = %v\n", n, result)
	}
}
