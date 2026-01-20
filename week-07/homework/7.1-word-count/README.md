# ДЗ 7.1: Подсчёт слов

## Цель
Научиться использовать map для подсчёта элементов.

## Что нужно сделать

1. Написать функцию `WordCount` — подсчитывает количество каждого слова
2. Написать функцию `CharCount` — подсчитывает количество каждого символа
3. Написать функцию `MostFrequent` — возвращает самое частое слово

## Сигнатуры

```go
func WordCount(text string) map[string]int
func CharCount(text string) map[rune]int
func MostFrequent(text string) string
```

## Пример использования

```go
func main() {
    text := "go go go python java go"

    words := WordCount(text)
    // map[go:4 python:1 java:1]

    chars := CharCount("hello")
    // map[h:1 e:1 l:2 o:1]

    most := MostFrequent(text)
    // "go"
}
```

## Подсказки

- Используй `strings.Fields(text)` для разбиения на слова
- Для подсчёта символов пройди по строке с `range`
- Для MostFrequent используй WordCount и найди максимум

## Критерии выполнения

- [ ] WordCount правильно считает все слова
- [ ] CharCount учитывает все символы включая пробелы
- [ ] MostFrequent возвращает слово с максимальным count
