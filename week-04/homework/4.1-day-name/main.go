package main

import "fmt"

// DayName возвращает название дня недели по номеру (1-7)
func DayName(n int) string {
	// TODO: реализуй с помощью switch
	switch n {
	case 1:
		return "Понедельник"
	// TODO: добавь остальные дни
	default:
		return "Некорректный день"
	}
}

func main() {
	for i := 0; i <= 8; i++ {
		fmt.Printf("День %d: %s\n", i, DayName(i))
	}
}
