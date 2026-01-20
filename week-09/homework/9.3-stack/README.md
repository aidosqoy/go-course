# ДЗ 9.3: Структура данных Stack

## Цель
Реализовать классическую структуру данных стек с помощью методов.

## Что нужно сделать

1. Создать структуру `Stack` со слайсом элементов
2. Написать методы:
   - `Push(value)` — добавляет элемент на вершину
   - `Pop()` — удаляет и возвращает верхний элемент
   - `Peek()` — возвращает верхний элемент без удаления
   - `IsEmpty()` — проверяет, пуст ли стек
   - `Size()` — возвращает количество элементов

## Сигнатуры

```go
type Stack struct {
    items []int
}

func NewStack() *Stack
func (s *Stack) Push(value int)
func (s *Stack) Pop() (int, bool)
func (s *Stack) Peek() (int, bool)
func (s *Stack) IsEmpty() bool
func (s *Stack) Size() int
```

## Пример использования

```go
func main() {
    stack := NewStack()

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println(stack.Size())    // 3
    fmt.Println(stack.Peek())    // 3, true
    fmt.Println(stack.Pop())     // 3, true
    fmt.Println(stack.Pop())     // 2, true
    fmt.Println(stack.Size())    // 1
    fmt.Println(stack.IsEmpty()) // false
}
```

## Подсказки

- LIFO — Last In, First Out (последний вошёл, первый вышел)
- Pop и Peek возвращают (value, ok) — ok=false если стек пуст

## Критерии выполнения

- [ ] Push добавляет на вершину
- [ ] Pop возвращает и удаляет верхний элемент
- [ ] Peek возвращает без удаления
- [ ] IsEmpty и Size работают корректно
- [ ] Pop и Peek безопасны для пустого стека
