package main

import "fmt"

func main() {
	// TODO: Объяви переменные всех базовых типов БЕЗ инициализации
	// и выведи их значения

	// Целые числа
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Println("=== Целые числа (signed) ===")
	fmt.Printf("int: %d\n", i)
	fmt.Printf("int8: %d\n", i8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("int64: %d\n", i64)

	// TODO: Добавь остальные типы:
	// - uint, uint8, uint16, uint32, uint64
	// - float32, float64
	// - bool
	// - string

	// TODO: Напиши ответ на вопрос здесь:
	// Почему в Go нет null/nil для базовых типов? Какие проблемы это решает?
	//
	// Ответ: ...
}
