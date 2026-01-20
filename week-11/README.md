# Неделя 11: Stringer, Error, Reader

## Теория

**Ссылка:** https://www.gocat.dev/tour/

Пройди уроки 75-83:
75. Stringers
76. Упражнение: Stringers
77. Ошибки
78. Упражнение: Ошибки
79. Readers
80. Упражнение: Readers
81. Упражнение: rot13Reader
82. Изображения
83. Упражнение: Изображения

## Домашние задания

| # | Задание | Описание |
|---|---------|----------|
| 11.1 | [custom-stringer](./homework/11.1-custom-stringer/) | IPAddr, Duration с String() |
| 11.2 | [custom-error](./homework/11.2-custom-error/) | ValidationError, NotFoundError |
| 11.3 | [string-reader](./homework/11.3-string-reader/) | RepeatReader, LimitReader |
| 11.4 | [rot13-reader](./homework/11.4-rot13-reader/) | ROT13 decoder (io.Reader wrapper) |

## Вопросы для самопроверки

Создай файл `answers-11.txt` и напиши ответы на вопросы:

1. Для чего используется интерфейс fmt.Stringer?
2. Как создать собственный тип ошибки в Go?
3. Что возвращает метод Read() интерфейса io.Reader?
4. Что означает io.EOF и когда его возвращать?

## Дополнительные материалы

- [Go by Example: Errors](https://gobyexample.com/errors)
- [Go by Example: Reading Files](https://gobyexample.com/reading-files)
- [Go Blog: Error handling](https://go.dev/blog/error-handling-and-go)
