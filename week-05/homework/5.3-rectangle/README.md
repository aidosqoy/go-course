# ДЗ 5.3: Структура Rectangle

## Цель
Подготовиться к изучению методов (пока через функции).

## Что нужно сделать

1. Создать структуру `Rectangle` с полями Width и Height (float64)
2. Написать функцию `Area` — вычисляет площадь
3. Написать функцию `Perimeter` — вычисляет периметр
4. Написать функцию `Scale` — масштабирует прямоугольник (через указатель)

## Сигнатуры

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func Area(r Rectangle) float64
func Perimeter(r Rectangle) float64
func Scale(r *Rectangle, factor float64)
```

## Пример использования

```go
func main() {
    rect := Rectangle{Width: 10, Height: 5}

    fmt.Printf("Площадь: %.2f\n", Area(rect))       // 50.00
    fmt.Printf("Периметр: %.2f\n", Perimeter(rect)) // 30.00

    Scale(&rect, 2)
    fmt.Printf("После масштабирования: %+v\n", rect) // {Width:20 Height:10}
}
```

## Формулы

- Площадь: `S = width × height`
- Периметр: `P = 2 × (width + height)`

## Критерии выполнения

- [ ] Структура Rectangle создана
- [ ] Формулы площади и периметра корректны
- [ ] Scale изменяет оба измерения через указатель
