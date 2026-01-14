package main

import "fmt"

/*
TODO: Напиши своё объяснение defer здесь

1. Что такое defer?


2. В каком порядке выполняются несколько defer?


3. Примеры использования defer в реальном коде:
   - Пример 1:
   - Пример 2:
   - Пример 3:

*/

func main() {
	// Демонстрация работы defer
	fmt.Println("Начало функции main")

	defer fmt.Println("Это выполнится последним (defer 1)")
	defer fmt.Println("Это выполнится предпоследним (defer 2)")
	defer fmt.Println("Это выполнится третьим с конца (defer 3)")

	fmt.Println("Конец функции main (но до defer)")
}
