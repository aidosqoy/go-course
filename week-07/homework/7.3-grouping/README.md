# ДЗ 7.3: Группировка данных

## Цель
Научиться использовать map со слайсами для группировки данных.

## Что нужно сделать

1. Создать структуру `Student` с полями Name и Grade (класс/курс)
2. Написать функцию `GroupByGrade` — группирует студентов по классу
3. Написать функцию `GroupByFirstLetter` — группирует слова по первой букве
4. Написать функцию `GetStudentNames` — возвращает имена студентов указанного класса

## Сигнатуры

```go
type Student struct {
    Name  string
    Grade int
}

func GroupByGrade(students []Student) map[int][]Student
func GroupByFirstLetter(words []string) map[rune][]string
func GetStudentNames(groups map[int][]Student, grade int) []string
```

## Пример использования

```go
func main() {
    students := []Student{
        {"Алексей", 10},
        {"Мария", 10},
        {"Иван", 11},
        {"Анна", 11},
    }

    groups := GroupByGrade(students)
    // map[10:[{Алексей 10} {Мария 10}] 11:[{Иван 11} {Анна 11}]]

    names := GetStudentNames(groups, 10)
    // ["Алексей", "Мария"]

    words := []string{"apple", "apricot", "banana", "avocado"}
    byLetter := GroupByFirstLetter(words)
    // map[a:[apple apricot avocado] b:[banana]]
}
```

## Критерии выполнения

- [ ] Student структура создана правильно
- [ ] GroupByGrade группирует по классу
- [ ] GroupByFirstLetter группирует по первой букве
- [ ] GetStudentNames возвращает только имена
