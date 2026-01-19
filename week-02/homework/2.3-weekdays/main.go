package main

import "fmt"

// TODO: Определи константы для дней недели с помощью iota
// Понедельник = 1, ..., Воскресенье = 7
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	// TODO: добавь остальные дни
	Thursday
	Friday
	Saturday
	Sunday
)

// IsWeekend возвращает true, если день является выходным (суббота или воскресенье)
func IsWeekend(day int) bool {
	// TODO: реализуй функцию

	if day == 6 || day == 7 {
		return true
	}

	return false
}

func main() {
	fmt.Println("Понедельник:", Monday)
	fmt.Println("Вторник:", Tuesday)
	fmt.Println("Среда:", Wednesday)
	// TODO: выведи остальные дни
	fmt.Println("Четверг:", Thursday)
	fmt.Println("Пятница:", Friday)
	fmt.Println("Суббота:", Saturday)
	fmt.Println("Воскресенье:", Sunday)

	fmt.Println()
	fmt.Println("Понедельник - выходной?", IsWeekend(Monday))
	fmt.Println("Среда - выходной?", IsWeekend(Wednesday))
	// TODO: проверь субботу и воскресенье
	fmt.Println("Суббота - выходной?", IsWeekend(Saturday))
	fmt.Println("Воскресенье - выходной?", IsWeekend(Sunday))

}
