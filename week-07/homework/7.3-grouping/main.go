package main

import "fmt"

// Student представляет информацию о студенте
type Student struct {
	Name  string
	Grade int
}

// GroupByGrade группирует студентов по классу
func GroupByGrade(students []Student) map[int][]Student {
	// TODO: реализуй функцию
	// Создай map[int][]Student
	// Пройди по студентам и добавляй в соответствующий слайс
	m := make(map[int][]Student)
	for _, v := range students {
		m[v.Grade] = append(m[v.Grade], v)
	}

	return m
}

// GroupByFirstLetter группирует слова по первой букве
func GroupByFirstLetter(words []string) map[rune][]string {
	// TODO: реализуй функцию
	// Для получения первой буквы используй []rune(word)[0]
	m := make(map[rune][]string)
	for _, v := range words {
		firstLetter := []rune(v)[0]
		m[firstLetter] = append(m[firstLetter], v)
	}
	return m
}

// GetStudentNames возвращает имена студентов указанного класса
func GetStudentNames(groups map[int][]Student, grade int) []string {
	// TODO: реализуй функцию
	// Получи слайс студентов по grade
	// Извлеки только имена
	students := groups[grade]

	var names []string
	for _, v := range students {
		names = append(names, v.Name)
	}
	return names
}

func main() {
	students := []Student{
		{"Алексей", 10},
		{"Мария", 10},
		{"Иван", 11},
		{"Анна", 11},
		{"Пётр", 10},
	}

	fmt.Println("=== Группировка студентов ===")
	fmt.Println("Все студенты:", students)

	groups := GroupByGrade(students)
	fmt.Println("\nПо классам:")
	for grade, list := range groups {
		fmt.Printf("  %d класс: %v\n", grade, list)
	}

	fmt.Println("\nИмена 10 класса:", GetStudentNames(groups, 10))

	fmt.Println("\n=== Группировка слов ===")
	words := []string{"apple", "apricot", "banana", "avocado", "blueberry"}
	byLetter := GroupByFirstLetter(words)
	fmt.Println("Слова:", words)
	fmt.Println("По первой букве:", byLetter)
}
