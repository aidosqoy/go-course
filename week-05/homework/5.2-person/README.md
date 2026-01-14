# ДЗ 5.2: Структура Person

## Цель
Научиться создавать и использовать структуры.

## Что нужно сделать

1. Создать структуру `Person` с полями: Name (string), Age (int), Email (string)
2. Написать функцию `NewPerson` — создаёт и возвращает Person
3. Написать функцию `PrintPerson` — красиво выводит информацию
4. Написать функцию `Birthday` — увеличивает возраст на 1 (через указатель)

## Сигнатуры

```go
type Person struct {
    Name  string
    Age   int
    Email string
}

func NewPerson(name string, age int, email string) Person
func PrintPerson(p Person)
func Birthday(p *Person)
```

## Пример использования

```go
func main() {
    person := NewPerson("Алексей", 25, "alex@example.com")
    PrintPerson(person)
    // Вывод:
    // Имя: Алексей
    // Возраст: 25
    // Email: alex@example.com

    Birthday(&person)
    fmt.Println("После дня рождения:", person.Age) // 26
}
```

## Критерии выполнения

- [ ] Структура Person создана правильно
- [ ] NewPerson возвращает заполненную структуру
- [ ] PrintPerson выводит читаемую информацию
- [ ] Birthday изменяет возраст через указатель
