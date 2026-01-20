# ДЗ 8.2: Замыкания — счётчик и аккумулятор

## Цель
Научиться использовать замыкания для сохранения состояния.

## Что нужно сделать

1. Написать функцию `Counter` — возвращает функцию-счётчик
2. Написать функцию `Accumulator` — возвращает функцию-аккумулятор
3. Написать функцию `Fibonacci` — возвращает функцию для генерации чисел Фибоначчи

## Сигнатуры

```go
func Counter() func() int
func Accumulator(initial int) func(int) int
func Fibonacci() func() int
```

## Пример использования

```go
func main() {
    // Counter
    count := Counter()
    fmt.Println(count()) // 1
    fmt.Println(count()) // 2
    fmt.Println(count()) // 3

    // Accumulator
    acc := Accumulator(10)
    fmt.Println(acc(5))  // 15
    fmt.Println(acc(3))  // 18
    fmt.Println(acc(-8)) // 10

    // Fibonacci
    fib := Fibonacci()
    fmt.Println(fib()) // 0
    fmt.Println(fib()) // 1
    fmt.Println(fib()) // 1
    fmt.Println(fib()) // 2
    fmt.Println(fib()) // 3
    fmt.Println(fib()) // 5
}
```

## Подсказки

- В замыкании переменная живёт между вызовами
- Fibonacci: храни два предыдущих числа, каждый раз вычисляй следующее

## Критерии выполнения

- [ ] Counter увеличивается при каждом вызове
- [ ] Accumulator накапливает сумму
- [ ] Fibonacci генерирует правильную последовательность
- [ ] Каждый вызов Counter/Accumulator создаёт независимый счётчик
