# ДЗ 7.4: Простой кэш

## Цель
Научиться создавать структуры с map для хранения данных.

## Что нужно сделать

1. Создать структуру `Cache` с полем data (map[string]string)
2. Написать функцию `NewCache` — создаёт новый кэш
3. Написать метод `Set` — сохраняет значение по ключу
4. Написать метод `Get` — получает значение по ключу
5. Написать метод `Delete` — удаляет значение
6. Написать метод `Clear` — очищает весь кэш
7. Написать метод `Size` — возвращает количество элементов

## Сигнатуры

```go
type Cache struct {
    data map[string]string
}

func NewCache() *Cache
func (c *Cache) Set(key, value string)
func (c *Cache) Get(key string) (string, bool)
func (c *Cache) Delete(key string)
func (c *Cache) Clear()
func (c *Cache) Size() int
```

## Пример использования

```go
func main() {
    cache := NewCache()

    cache.Set("user:1", "Алексей")
    cache.Set("user:2", "Мария")

    name, ok := cache.Get("user:1")
    // "Алексей", true

    size := cache.Size() // 2

    cache.Delete("user:1")
    cache.Clear()
}
```

## Критерии выполнения

- [ ] Cache структура создана с приватным полем data
- [ ] NewCache инициализирует map
- [ ] Set сохраняет пару ключ-значение
- [ ] Get возвращает значение и флаг существования
- [ ] Delete удаляет по ключу
- [ ] Clear очищает весь кэш
- [ ] Size возвращает количество элементов
