# ДЗ 9.2: Банковский счёт

## Цель
Научиться использовать методы для изменения состояния структуры.

## Что нужно сделать

1. Создать структуру `Account` с полями Owner и Balance
2. Написать методы:
   - `Deposit(amount)` — пополнение счёта
   - `Withdraw(amount)` — снятие (с проверкой баланса)
   - `Transfer(to *Account, amount)` — перевод на другой счёт
   - `Statement()` — выписка (строка с информацией)

## Сигнатуры

```go
type Account struct {
    Owner   string
    Balance float64
}

func NewAccount(owner string, initial float64) *Account
func (a *Account) Deposit(amount float64) error
func (a *Account) Withdraw(amount float64) error
func (a *Account) Transfer(to *Account, amount float64) error
func (a Account) Statement() string
```

## Пример использования

```go
func main() {
    acc1 := NewAccount("Алексей", 1000)
    acc2 := NewAccount("Мария", 500)

    acc1.Deposit(500)        // Баланс: 1500
    acc1.Withdraw(200)       // Баланс: 1300
    acc1.Transfer(acc2, 300) // acc1: 1000, acc2: 800

    err := acc1.Withdraw(5000) // Ошибка: недостаточно средств

    fmt.Println(acc1.Statement())
    // "Счёт: Алексей, Баланс: 1000.00"
}
```

## Критерии выполнения

- [ ] Deposit увеличивает баланс (проверка amount > 0)
- [ ] Withdraw уменьшает баланс (проверка достаточности средств)
- [ ] Transfer переводит между счетами атомарно
- [ ] Все методы, изменяющие состояние, используют указатели
- [ ] Ошибки возвращаются для невалидных операций
