# ДЗ 8.3: Калькулятор с map операций

## Цель
Научиться хранить функции в map и вызывать их по ключу.

## Что нужно сделать

1. Создать тип `Operation` — функция (a, b int) int
2. Создать map операций с ключами "+", "-", "*", "/"
3. Написать функцию `Calculate` — выполняет операцию по символу
4. Написать функцию `RegisterOperation` — добавляет новую операцию

## Сигнатуры

```go
type Operation func(a, b int) int

var operations map[string]Operation

func Calculate(op string, a, b int) (int, error)
func RegisterOperation(op string, fn Operation)
```

## Пример использования

```go
func main() {
    result, _ := Calculate("+", 10, 5)  // 15
    result, _ = Calculate("-", 10, 5)   // 5
    result, _ = Calculate("*", 10, 5)   // 50
    result, _ = Calculate("/", 10, 5)   // 2

    // Добавляем новую операцию — возведение в степень
    RegisterOperation("^", func(a, b int) int {
        result := 1
        for i := 0; i < b; i++ {
            result *= a
        }
        return result
    })

    result, _ = Calculate("^", 2, 8)    // 256
}
```

## Критерии выполнения

- [ ] Базовые операции +, -, *, / работают
- [ ] Деление на ноль обрабатывается (возврат ошибки)
- [ ] RegisterOperation добавляет новые операции
- [ ] Calculate возвращает ошибку для неизвестной операции
