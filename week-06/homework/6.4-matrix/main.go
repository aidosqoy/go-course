package main

import "fmt"

// CreateMatrix создаёт матрицу размером rows×cols, заполненную значением value
func CreateMatrix(rows, cols, value int) [][]int {
	// TODO: реализуй функцию
	// Создай слайс слайсов нужного размера
	// Заполни каждый элемент значением value
	return nil
}

// PrintMatrix выводит матрицу в читаемом формате
func PrintMatrix(matrix [][]int) {
	// TODO: реализуй функцию
	// Выведи каждую строку на отдельной линии
	// Элементы разделяй пробелами или табуляцией
	fmt.Println(matrix)
}

// Transpose возвращает транспонированную матрицу
func Transpose(matrix [][]int) [][]int {
	// TODO: реализуй функцию
	// Если исходная матрица M×N, результат будет N×M
	// Элемент [i][j] становится [j][i]
	return nil
}

// SumMatrix возвращает сумму всех элементов матрицы
func SumMatrix(matrix [][]int) int {
	// TODO: реализуй функцию
	// Пройди по всем строкам и столбцам
	return 0
}

func main() {
	fmt.Println("=== Работа с матрицами ===")

	// Создание матрицы
	matrix := CreateMatrix(2, 3, 1)
	fmt.Println("Созданная матрица 2×3:")
	PrintMatrix(matrix)

	// Работа с готовой матрицей
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\nИсходная матрица:")
	PrintMatrix(matrix)

	fmt.Println("\nТранспонированная матрица:")
	transposed := Transpose(matrix)
	PrintMatrix(transposed)

	fmt.Println("\nСумма элементов:", SumMatrix(matrix))
}
