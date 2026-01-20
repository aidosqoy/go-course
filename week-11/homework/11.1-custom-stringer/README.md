# ДЗ 11.1: Реализация Stringer

## Цель
Научиться реализовывать интерфейс `fmt.Stringer` для красивого вывода.

## Что нужно сделать

1. Создать тип `IPAddr` ([4]byte) и реализовать String()
2. Создать тип `Duration` (int — секунды) и реализовать String()
3. Создать структуру `Book` и реализовать String()

## Сигнатуры

```go
type IPAddr [4]byte

func (ip IPAddr) String() string

type Duration int

func (d Duration) String() string

type Book struct {
    Title  string
    Author string
    Year   int
}

func (b Book) String() string
```

## fmt.Stringer

```go
type Stringer interface {
    String() string
}
```

## Пример использования

```go
func main() {
    ip := IPAddr{192, 168, 0, 1}
    fmt.Println(ip) // "192.168.0.1"

    d := Duration(3665) // 1 час, 1 минута, 5 секунд
    fmt.Println(d)      // "1h 1m 5s"

    book := Book{"Война и мир", "Лев Толстой", 1869}
    fmt.Println(book)   // "«Война и мир» (Лев Толстой, 1869)"
}
```

## Подсказки

- Для IP используй `fmt.Sprintf("%d.%d.%d.%d", ...)`
- Для Duration раздели на часы, минуты, секунды

## Критерии выполнения

- [ ] IPAddr выводит в формате "a.b.c.d"
- [ ] Duration выводит в формате "Xh Xm Xs"
- [ ] Book выводит название, автора и год
- [ ] fmt.Println использует String() автоматически
