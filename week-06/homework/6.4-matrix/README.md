# ДЗ 6.4: Двумерные массивы (матрицы)

## Цель
Научиться работать с двумерными массивами и слайсами.

## Что нужно сделать

1. Написать функцию `CreateMatrix` — создаёт матрицу rows×cols заполненную value
2. Написать функцию `PrintMatrix` — красиво выводит матрицу
3. Написать функцию `Transpose` — транспонирует матрицу (строки ↔ столбцы)
4. Написать функцию `SumMatrix` — возвращает сумму всех элементов

## Сигнатуры

```go
func CreateMatrix(rows, cols, value int) [][]int
func PrintMatrix(matrix [][]int)
func Transpose(matrix [][]int) [][]int
func SumMatrix(matrix [][]int) int
```

## Пример использования

```go
func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    PrintMatrix(matrix)
    // 1 2 3
    // 4 5 6

    transposed := Transpose(matrix)
    PrintMatrix(transposed)
    // 1 4
    // 2 5
    // 3 6

    sum := SumMatrix(matrix) // 21

    newMatrix := CreateMatrix(2, 3, 0)
    // [[0, 0, 0], [0, 0, 0]]
}
```

## Подсказки

- При транспонировании матрица M×N становится N×M
- Элемент [i][j] становится [j][i]

## Критерии выполнения

- [ ] CreateMatrix создаёт матрицу нужного размера
- [ ] PrintMatrix выводит читаемый формат
- [ ] Transpose корректно меняет строки и столбцы
- [ ] SumMatrix правильно суммирует все элементы
