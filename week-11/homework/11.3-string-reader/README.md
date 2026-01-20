# ДЗ 11.3: Реализация io.Reader

## Цель
Научиться реализовывать интерфейс io.Reader для создания источников данных.

## Что нужно сделать

1. Создать `RepeatReader` — бесконечно повторяет один байт
2. Создать `LimitReader` — ограничивает количество байт из другого Reader
3. Создать `CountingReader` — считает прочитанные байты

## Сигнатуры

```go
type RepeatReader struct {
    char byte
}

func NewRepeatReader(char byte) *RepeatReader
func (r *RepeatReader) Read(p []byte) (n int, err error)

type LimitReader struct {
    reader io.Reader
    limit  int
    read   int
}

func NewLimitReader(r io.Reader, limit int) *LimitReader
func (r *LimitReader) Read(p []byte) (n int, err error)

type CountingReader struct {
    reader    io.Reader
    BytesRead int
}

func NewCountingReader(r io.Reader) *CountingReader
func (r *CountingReader) Read(p []byte) (n int, err error)
```

## io.Reader

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## Пример использования

```go
func main() {
    // RepeatReader — бесконечно выдаёт 'A'
    rr := NewRepeatReader('A')
    buf := make([]byte, 5)
    rr.Read(buf)
    fmt.Println(string(buf)) // "AAAAA"

    // LimitReader — читает максимум n байт
    original := strings.NewReader("Hello, World!")
    lr := NewLimitReader(original, 5)
    data, _ := io.ReadAll(lr)
    fmt.Println(string(data)) // "Hello"

    // CountingReader — считает байты
    cr := NewCountingReader(strings.NewReader("test"))
    io.ReadAll(cr)
    fmt.Println(cr.BytesRead) // 4
}
```

## Критерии выполнения

- [ ] RepeatReader бесконечно генерирует один символ
- [ ] LimitReader останавливается после limit байт
- [ ] CountingReader правильно считает прочитанные байты
- [ ] Все реализации корректно возвращают io.EOF
