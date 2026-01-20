# ДЗ 9.1: Геометрические фигуры

## Цель
Научиться создавать методы для структур.

## Что нужно сделать

1. Создать структуру `Circle` с полем Radius
2. Создать структуру `Rectangle` с полями Width и Height
3. Для каждой фигуры реализовать методы:
   - `Area()` — площадь
   - `Perimeter()` — периметр
   - `Scale(factor float64)` — масштабирование

## Сигнатуры

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64
func (c Circle) Perimeter() float64
func (c *Circle) Scale(factor float64)

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64
func (r Rectangle) Perimeter() float64
func (r *Rectangle) Scale(factor float64)
```

## Пример использования

```go
func main() {
    c := Circle{Radius: 5}
    fmt.Println("Площадь круга:", c.Area())         // ~78.54
    fmt.Println("Периметр круга:", c.Perimeter())   // ~31.42
    c.Scale(2)
    fmt.Println("После Scale(2), радиус:", c.Radius) // 10

    r := Rectangle{Width: 4, Height: 3}
    fmt.Println("Площадь прямоугольника:", r.Area())       // 12
    fmt.Println("Периметр прямоугольника:", r.Perimeter()) // 14
}
```

## Подсказки

- Используй `math.Pi` для числа π
- Scale должен изменять оригинал — используй указатель

## Критерии выполнения

- [ ] Circle.Area() вычисляет π × r²
- [ ] Circle.Perimeter() вычисляет 2 × π × r
- [ ] Rectangle.Area() вычисляет width × height
- [ ] Rectangle.Perimeter() вычисляет 2 × (width + height)
- [ ] Scale изменяет размеры через указатель
