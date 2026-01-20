# ДЗ 10.4: Реализация sort.Interface

## Цель
Научиться реализовывать стандартные интерфейсы Go на примере sort.Interface.

## Что нужно сделать

1. Создать структуру `Person` с полями Name и Age
2. Создать тип `ByAge` (слайс Person) и реализовать sort.Interface
3. Создать тип `ByName` (слайс Person) и реализовать sort.Interface
4. Продемонстрировать сортировку по разным полям

## Сигнатуры

```go
type Person struct {
    Name string
    Age  int
}

type ByAge []Person

func (a ByAge) Len() int
func (a ByAge) Swap(i, j int)
func (a ByAge) Less(i, j int) bool

type ByName []Person

func (a ByName) Len() int
func (a ByName) Swap(i, j int)
func (a ByName) Less(i, j int) bool
```

## sort.Interface

```go
type Interface interface {
    Len() int           // количество элементов
    Less(i, j int) bool // true если элемент i должен быть перед j
    Swap(i, j int)      // меняет элементы местами
}
```

## Пример использования

```go
func main() {
    people := []Person{
        {"Алексей", 25},
        {"Мария", 30},
        {"Иван", 20},
    }

    sort.Sort(ByAge(people))
    // [{Иван 20} {Алексей 25} {Мария 30}]

    sort.Sort(ByName(people))
    // [{Алексей 25} {Иван 20} {Мария 30}]
}
```

## Критерии выполнения

- [ ] Person структура создана
- [ ] ByAge реализует sort.Interface
- [ ] ByName реализует sort.Interface
- [ ] sort.Sort() работает с обоими типами
