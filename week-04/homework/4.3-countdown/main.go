package main

import "fmt"

// Countdown выводит обратный отсчёт от n до 1, затем "Поехали!"
func Countdown(n int) {
	// TODO: реализуй функцию используя defer
	defer fmt.Println("Поехали!")

	for i := n; i >= 1; i-- {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("=== Обратный отсчёт ===")
	Countdown(5)
}
