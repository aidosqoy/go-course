# ДЗ 7.2: Телефонная книга

## Цель
Научиться выполнять CRUD операции с map.

## Что нужно сделать

1. Написать функцию `AddContact` — добавляет контакт
2. Написать функцию `GetContact` — получает номер по имени
3. Написать функцию `UpdateContact` — обновляет номер
4. Написать функцию `DeleteContact` — удаляет контакт
5. Написать функцию `ListContacts` — выводит все контакты

## Сигнатуры

```go
func AddContact(book map[string]string, name, phone string)
func GetContact(book map[string]string, name string) (string, bool)
func UpdateContact(book map[string]string, name, phone string) bool
func DeleteContact(book map[string]string, name string) bool
func ListContacts(book map[string]string)
```

## Пример использования

```go
func main() {
    book := make(map[string]string)

    AddContact(book, "Алексей", "+7-999-123-4567")
    AddContact(book, "Мария", "+7-999-765-4321")

    phone, ok := GetContact(book, "Алексей")
    // "+7-999-123-4567", true

    UpdateContact(book, "Алексей", "+7-999-000-0000")
    DeleteContact(book, "Мария")

    ListContacts(book)
    // Алексей: +7-999-000-0000
}
```

## Критерии выполнения

- [ ] AddContact добавляет новый контакт
- [ ] GetContact возвращает номер и флаг существования
- [ ] UpdateContact обновляет только существующий контакт
- [ ] DeleteContact удаляет контакт, если он есть
- [ ] ListContacts выводит все контакты
