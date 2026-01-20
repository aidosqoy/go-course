# ДЗ 9.4: Встраивание структур

## Цель
Научиться использовать встраивание (embedding) для расширения структур.

## Что нужно сделать

1. Создать структуру `Person` с полями Name и Age
2. Создать структуру `Employee` — встраивает Person, добавляет Position и Salary
3. Создать структуру `Manager` — встраивает Employee, добавляет Team (слайс имён)
4. Реализовать методы для каждого уровня иерархии

## Сигнатуры

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string

type Employee struct {
    Person
    Position string
    Salary   float64
}

func (e Employee) Work() string
func (e *Employee) GiveRaise(percent float64)

type Manager struct {
    Employee
    Team []string
}

func (m *Manager) AddToTeam(name string)
func (m Manager) TeamSize() int
```

## Пример использования

```go
func main() {
    manager := Manager{
        Employee: Employee{
            Person:   Person{Name: "Алексей", Age: 35},
            Position: "Tech Lead",
            Salary:   150000,
        },
        Team: []string{},
    }

    // Методы Person доступны напрямую
    fmt.Println(manager.Greet())      // "Привет, меня зовут Алексей"
    fmt.Println(manager.Name)         // "Алексей" (поле Person)

    // Методы Employee
    fmt.Println(manager.Work())       // "Алексей работает как Tech Lead"
    manager.GiveRaise(10)             // Повышение зарплаты на 10%

    // Методы Manager
    manager.AddToTeam("Мария")
    manager.AddToTeam("Иван")
    fmt.Println(manager.TeamSize())   // 2
}
```

## Критерии выполнения

- [ ] Person.Greet() возвращает приветствие с именем
- [ ] Employee встраивает Person (поля и методы доступны)
- [ ] Employee.Work() использует поля из Person
- [ ] Manager встраивает Employee
- [ ] Методы работают через всю иерархию
