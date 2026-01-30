package main

import "fmt"

// CreateMatrix создаёт матрицу размером rows×cols, заполненную значением value
func CreateMatrix(rows, cols, value int) [][]int {
	// TODO: реализуй функцию
	// Создай слайс слайсов нужного размера
	// Заполни каждый элемент значением value

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = value
		}
	}

	return matrix
}

// PrintMatrix выводит матрицу в читаемом формате
func PrintMatrix(matrix [][]int) {
	// TODO: реализуй функцию
	// Выведи каждую строку на отдельной линии
	// Элементы разделяй пробелами или табуляцией
	for _, row := range matrix {
		for i, value := range row {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
}

// Transpose возвращает транспонированную матрицу
func Transpose(matrix [][]int) [][]int {
	// TODO: реализуй функцию
	// Если исходная матрица M×N, результат будет N×M
	// Элемент [i][j] становится [j][i]
	if len(matrix) == 0 {
		return [][]int{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// SumMatrix возвращает сумму всех элементов матрицы
func SumMatrix(matrix [][]int) int {
	// TODO: реализуй функцию
	// Пройди по всем строкам и столбцам
	sum := 0

	for _, row := range matrix {
		for _, value := range row {
			sum += value
		}
	}

	return sum
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
