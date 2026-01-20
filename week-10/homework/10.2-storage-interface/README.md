# ДЗ 10.2: Интерфейс Storage

## Цель
Научиться использовать интерфейсы для создания взаимозаменяемых реализаций.

## Что нужно сделать

1. Создать интерфейс `Storage` с методами Set, Get, Delete, Keys
2. Реализовать `MemoryStorage` — хранение в map
3. Реализовать `SliceStorage` — хранение в слайсе пар ключ-значение
4. Написать функцию `CopyStorage` — копирует данные между любыми Storage

## Сигнатуры

```go
type Storage interface {
    Set(key, value string)
    Get(key string) (string, bool)
    Delete(key string)
    Keys() []string
}

type MemoryStorage struct {
    data map[string]string
}

type SliceStorage struct {
    items []KeyValue
}

type KeyValue struct {
    Key, Value string
}

func NewMemoryStorage() *MemoryStorage
func NewSliceStorage() *SliceStorage
func CopyStorage(from, to Storage)
```

## Пример использования

```go
func main() {
    mem := NewMemoryStorage()
    mem.Set("name", "Алексей")
    mem.Set("city", "Москва")

    slice := NewSliceStorage()
    CopyStorage(mem, slice) // Копируем из mem в slice

    val, _ := slice.Get("name") // "Алексей"
    fmt.Println(slice.Keys())   // ["name", "city"]
}
```

## Критерии выполнения

- [ ] Storage интерфейс определён правильно
- [ ] MemoryStorage реализует все методы
- [ ] SliceStorage реализует все методы
- [ ] CopyStorage работает с любыми реализациями Storage
