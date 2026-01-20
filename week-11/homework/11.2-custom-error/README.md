# ДЗ 11.2: Пользовательские ошибки

## Цель
Научиться создавать собственные типы ошибок, реализующие интерфейс error.

## Что нужно сделать

1. Создать тип `ValidationError` — ошибка валидации с полем и сообщением
2. Создать тип `NotFoundError` — ошибка "не найдено" с типом и ID
3. Написать функции, возвращающие эти ошибки
4. Использовать errors.As для проверки типа ошибки

## Сигнатуры

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string

type NotFoundError struct {
    Resource string
    ID       int
}

func (e NotFoundError) Error() string

func ValidateAge(age int) error
func FindUser(id int) (string, error)
```

## error interface

```go
type error interface {
    Error() string
}
```

## Пример использования

```go
func main() {
    err := ValidateAge(-5)
    if err != nil {
        fmt.Println(err) // "поле 'age': возраст не может быть отрицательным"

        var valErr ValidationError
        if errors.As(err, &valErr) {
            fmt.Println("Поле:", valErr.Field)
        }
    }

    _, err = FindUser(999)
    if err != nil {
        fmt.Println(err) // "user с ID 999 не найден"

        var notFound NotFoundError
        if errors.As(err, &notFound) {
            fmt.Println("Ресурс:", notFound.Resource)
        }
    }
}
```

## Критерии выполнения

- [ ] ValidationError реализует интерфейс error
- [ ] NotFoundError реализует интерфейс error
- [ ] errors.As позволяет извлечь конкретный тип ошибки
- [ ] Error() возвращает информативное сообщение
