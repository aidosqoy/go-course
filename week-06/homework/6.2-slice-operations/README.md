# ДЗ 6.2: Операции со слайсами

## Цель
Научиться работать со слайсами: добавление, удаление, variadic функции.

## Что нужно сделать

1. Написать функцию `AppendUnique` — добавляет элемент только если его нет в слайсе
2. Написать функцию `RemoveAt` — удаляет элемент по индексу
3. Написать функцию `RemoveValue` — удаляет первое вхождение значения
4. Написать функцию `SumAll` — variadic функция для суммы любого количества чисел

## Сигнатуры

```go
func AppendUnique(slice []int, value int) []int
func RemoveAt(slice []int, index int) []int
func RemoveValue(slice []int, value int) []int
func SumAll(nums ...int) int
```

## Пример использования

```go
func main() {
    nums := []int{1, 2, 3, 4, 5}

    nums = AppendUnique(nums, 6)  // [1, 2, 3, 4, 5, 6]
    nums = AppendUnique(nums, 3)  // [1, 2, 3, 4, 5, 6] (3 уже есть)

    nums = RemoveAt(nums, 2)      // [1, 2, 4, 5, 6]
    nums = RemoveValue(nums, 5)   // [1, 2, 4, 6]

    total := SumAll(1, 2, 3, 4, 5) // 15
    total = SumAll(nums...)        // передача слайса
}
```

## Критерии выполнения

- [ ] AppendUnique не добавляет дубликаты
- [ ] RemoveAt корректно удаляет по индексу (обработай выход за границы)
- [ ] RemoveValue удаляет только первое вхождение
- [ ] SumAll работает с любым количеством аргументов
