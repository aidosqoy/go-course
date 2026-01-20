# ДЗ 10.1: Интерфейс Shape

## Цель
Научиться создавать и использовать интерфейсы для полиморфизма.

## Что нужно сделать

1. Создать интерфейс `Shape` с методами Area() и Perimeter()
2. Реализовать интерфейс для `Circle` и `Rectangle`
3. Написать функцию `PrintShapeInfo` — принимает любой Shape
4. Написать функцию `TotalArea` — суммирует площади слайса фигур

## Сигнатуры

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

func PrintShapeInfo(s Shape)
func TotalArea(shapes []Shape) float64
```

## Пример использования

```go
func main() {
    shapes := []Shape{
        Circle{Radius: 5},
        Rectangle{Width: 4, Height: 3},
        Circle{Radius: 2},
    }

    for _, s := range shapes {
        PrintShapeInfo(s)
    }
    // Вывод для каждой фигуры: площадь и периметр

    total := TotalArea(shapes)
    fmt.Println("Общая площадь:", total)
}
```

## Подсказки

- Интерфейс реализуется неявно — достаточно иметь нужные методы
- Для π используй `math.Pi`

## Критерии выполнения

- [ ] Интерфейс Shape определён с двумя методами
- [ ] Circle реализует Shape
- [ ] Rectangle реализует Shape
- [ ] PrintShapeInfo работает с любым Shape
- [ ] TotalArea суммирует площади всех фигур
