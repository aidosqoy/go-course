# ДЗ 8.4: Композиция функций (Middleware)

## Цель
Научиться составлять цепочки функций для обработки данных.

## Что нужно сделать

1. Создать тип `StringProcessor` — функция (string) string
2. Написать функции обработки: `ToUpper`, `ToLower`, `Trim`, `AddPrefix`
3. Написать функцию `Compose` — объединяет несколько процессоров в один
4. Написать функцию `Pipeline` — применяет процессоры последовательно

## Сигнатуры

```go
type StringProcessor func(string) string

func ToUpper() StringProcessor
func ToLower() StringProcessor
func Trim() StringProcessor
func AddPrefix(prefix string) StringProcessor

func Compose(processors ...StringProcessor) StringProcessor
func Pipeline(input string, processors ...StringProcessor) string
```

## Пример использования

```go
func main() {
    // Отдельные процессоры
    upper := ToUpper()
    result := upper("hello")  // "HELLO"

    // Compose — создаёт один процессор из нескольких
    processor := Compose(Trim(), ToLower(), AddPrefix(">>> "))
    result = processor("  HELLO WORLD  ")
    // ">>> hello world"

    // Pipeline — применяет процессоры к строке
    result = Pipeline("  HELLO  ",
        Trim(),
        ToLower(),
        AddPrefix("Result: "),
    )
    // "Result: hello"
}
```

## Подсказки

- Используй `strings.ToUpper`, `strings.ToLower`, `strings.TrimSpace`
- AddPrefix возвращает функцию (замыкание), которая запоминает prefix

## Критерии выполнения

- [ ] Базовые процессоры работают корректно
- [ ] AddPrefix использует замыкание для хранения prefix
- [ ] Compose объединяет процессоры в цепочку
- [ ] Pipeline применяет процессоры слева направо
