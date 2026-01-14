package main

import "fmt"

// FizzBuzz выводит числа от 1 до n по правилам FizzBuzz
func FizzBuzz(n int) {
	// TODO: реализуй функцию
	// Правила:
	// - Делится на 3 и 5 → "FizzBuzz"
	// - Делится на 3 → "Fizz"
	// - Делится на 5 → "Buzz"
	// - Иначе → само число
	for i := 1; i <= n; i++ {
		// TODO: добавь логику
		fmt.Println(i)
	}
}

func main() {
	FizzBuzz(15)
}
