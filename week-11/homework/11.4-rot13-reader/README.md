# ДЗ 11.4: ROT13 Reader

## Цель
Научиться создавать Reader-обёртки для трансформации данных.

## Что нужно сделать

1. Создать `Rot13Reader` — обёртка, которая применяет ROT13 к данным
2. ROT13: A→N, B→O, ..., M→Z, N→A, ..., Z→M (и аналогично для маленьких букв)
3. Написать вспомогательную функцию `rot13(b byte) byte`

## Сигнатуры

```go
type Rot13Reader struct {
    reader io.Reader
}

func NewRot13Reader(r io.Reader) *Rot13Reader
func (r *Rot13Reader) Read(p []byte) (n int, err error)
func rot13(b byte) byte
```

## ROT13

ROT13 — простой шифр замены, сдвигающий каждую букву на 13 позиций:
- A↔N, B↔O, C↔P, ..., M↔Z
- a↔n, b↔o, c↔p, ..., m↔z
- Цифры и другие символы не меняются

Особенность: ROT13(ROT13(x)) = x (двойное применение возвращает оригинал)

## Пример использования

```go
func main() {
    // Кодируем
    input := strings.NewReader("Hello, World!")
    r := NewRot13Reader(input)
    encoded, _ := io.ReadAll(r)
    fmt.Println(string(encoded)) // "Uryyb, Jbeyq!"

    // Декодируем (ROT13 симметричен)
    input2 := strings.NewReader("Uryyb, Jbeyq!")
    r2 := NewRot13Reader(input2)
    decoded, _ := io.ReadAll(r2)
    fmt.Println(string(decoded)) // "Hello, World!"
}
```

## Критерии выполнения

- [ ] rot13() корректно сдвигает буквы A-Z и a-z
- [ ] rot13() не меняет цифры и спецсимволы
- [ ] Rot13Reader реализует io.Reader
- [ ] Двойное применение возвращает оригинал
