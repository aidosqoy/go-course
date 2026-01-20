# ДЗ 8.1: Функции как значения

## Цель
Научиться передавать функции как аргументы (функции высшего порядка).

## Что нужно сделать

1. Написать функцию `Apply` — применяет функцию к каждому элементу слайса
2. Написать функцию `Filter` — фильтрует слайс по предикату
3. Написать функцию `Reduce` — сворачивает слайс в одно значение

## Сигнатуры

```go
func Apply(nums []int, fn func(int) int) []int
func Filter(nums []int, predicate func(int) bool) []int
func Reduce(nums []int, initial int, fn func(int, int) int) int
```

## Пример использования

```go
func main() {
    nums := []int{1, 2, 3, 4, 5}

    // Apply: умножаем каждый элемент на 2
    doubled := Apply(nums, func(x int) int { return x * 2 })
    // [2, 4, 6, 8, 10]

    // Filter: оставляем только чётные
    even := Filter(nums, func(x int) bool { return x%2 == 0 })
    // [2, 4]

    // Reduce: суммируем все элементы
    sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
    // 15
}
```

## Критерии выполнения

- [ ] Apply применяет функцию ко всем элементам
- [ ] Filter возвращает только элементы, для которых predicate = true
- [ ] Reduce корректно сворачивает слайс
- [ ] Оригинальный слайс не изменяется
