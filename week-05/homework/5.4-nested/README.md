# ДЗ 5.4: Вложенные структуры

## Цель
Научиться работать с композицией структур.

## Что нужно сделать

1. Создать структуру `Address` с полями: City, Street, Building (string)
2. Добавить поле Address в структуру Person
3. Написать функцию создания Person с полным адресом

## Структуры

```go
type Address struct {
    City     string
    Street   string
    Building string
}

type Person struct {
    Name    string
    Age     int
    Email   string
    Address Address  // вложенная структура
}
```

## Пример использования

```go
func main() {
    person := Person{
        Name:  "Мария",
        Age:   30,
        Email: "maria@example.com",
        Address: Address{
            City:     "Москва",
            Street:   "Тверская",
            Building: "1",
        },
    }

    fmt.Println(person.Name)            // Мария
    fmt.Println(person.Address.City)    // Москва
    fmt.Println(person.Address.Street)  // Тверская
}
```

## Критерии выполнения

- [ ] Структура Address создана
- [ ] Person содержит вложенный Address
- [ ] Доступ к полям вложенной структуры работает
