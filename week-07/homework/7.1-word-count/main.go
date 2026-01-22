package main

import (
	"fmt"
	"strings"
)

// WordCount возвращает map с количеством каждого слова в тексте
func WordCount(text string) map[string]int {
	// TODO: реализуй функцию
	// Используй strings.Fields для разбиения на слова
	// Пройди по словам и увеличивай счётчик в map

	txtSlice := strings.Fields(text)

	m := make(map[string]int)
	for _, v := range txtSlice {
		_, ok := m[v]
		if ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

// CharCount возвращает map с количеством каждого символа
func CharCount(text string) map[rune]int {
	// TODO: реализуй функцию
	// Пройди по строке с range — получишь rune
	m := make(map[rune]int)
	for _, v := range text {
		m[v] = m[v] + 1
	}

	return m
}

// MostFrequent возвращает самое часто встречающееся слово
func MostFrequent(text string) string {
	// TODO: реализуй функцию
	// Используй WordCount, затем найди слово с максимальным значением
	m := WordCount(text)

	maxCount := 0
	key := ""
	for k, v := range m {
		if v > maxCount {
			maxCount = v
			key = k
		}
	}

	return key
}

func main() {
	text := "go go go python java go python"

	fmt.Println("=== Подсчёт слов ===")
	fmt.Println("Текст:", text)
	fmt.Println("Количество слов:", WordCount(text))
	fmt.Println("Самое частое слово:", MostFrequent(text))

	fmt.Println("\n=== Подсчёт символов ===")
	word := "hello"
	fmt.Println("Слово:", word)
	fmt.Println("Количество символов:", CharCount(word))
}
