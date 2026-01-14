# ДЗ 2.2: Конвертер температуры

## Цель
Научиться создавать собственные типы и преобразовывать между ними.

## Что нужно сделать

1. Создать типы `Celsius` и `Fahrenheit` на основе float64
2. Написать функции конвертации между ними

## Формулы

- Цельсий → Фаренгейт: `F = C × 9/5 + 32`
- Фаренгейт → Цельсий: `C = (F - 32) × 5/9`

## Сигнатуры функций

```go
type Celsius float64
type Fahrenheit float64

func CelsiusToFahrenheit(c Celsius) Fahrenheit
func FahrenheitToCelsius(f Fahrenheit) Celsius
```

## Пример использования

```go
func main() {
    c := Celsius(100)
    f := CelsiusToFahrenheit(c)
    fmt.Printf("%.1f°C = %.1f°F\n", c, f) // 100.0°C = 212.0°F

    f2 := Fahrenheit(32)
    c2 := FahrenheitToCelsius(f2)
    fmt.Printf("%.1f°F = %.1f°C\n", f2, c2) // 32.0°F = 0.0°C
}
```

## Тестовые значения

| Цельсий | Фаренгейт |
|---------|-----------|
| 0       | 32        |
| 100     | 212       |
| -40     | -40       |
| 37      | 98.6      |

## Критерии выполнения

- [ ] Созданы кастомные типы Celsius и Fahrenheit
- [ ] Обе функции конвертации работают корректно
- [ ] Проверено на тестовых значениях из таблицы
