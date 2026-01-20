# ДЗ 10.3: Пустой интерфейс и type switch

## Цель
Научиться работать с `any` (пустой интерфейс), type assertion и type switch.

## Что нужно сделать

1. Написать функцию `TypeName` — возвращает название типа значения
2. Написать функцию `ToString` — преобразует любое значение в строку
3. Написать функцию `Sum` — суммирует числа из слайса any
4. Написать функцию `FilterByType` — фильтрует слайс по типу

## Сигнатуры

```go
func TypeName(v any) string
func ToString(v any) string
func Sum(values []any) float64
func FilterByType[T any](values []any) []T
```

## Пример использования

```go
func main() {
    fmt.Println(TypeName(42))        // "int"
    fmt.Println(TypeName("hello"))   // "string"
    fmt.Println(TypeName(3.14))      // "float64"

    fmt.Println(ToString(42))        // "42"
    fmt.Println(ToString(true))      // "true"
    fmt.Println(ToString([]int{1}))  // "[1]"

    values := []any{1, 2.5, "skip", 3, 4.5}
    fmt.Println(Sum(values))         // 11.0 (1 + 2.5 + 3 + 4.5)

    mixed := []any{1, "hello", 2, "world", 3}
    strings := FilterByType[string](mixed)
    // ["hello", "world"]
}
```

## Подсказки

- Используй type switch для определения типа
- Type assertion: `v.(int)` или `v.(string)`
- Для ToString используй `fmt.Sprintf("%v", v)`

## Критерии выполнения

- [ ] TypeName возвращает правильное название типа
- [ ] ToString работает с основными типами
- [ ] Sum суммирует только числа (int, float64)
- [ ] FilterByType возвращает только элементы нужного типа
