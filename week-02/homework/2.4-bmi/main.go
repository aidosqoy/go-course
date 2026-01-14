package main

import "fmt"

// CalculateBMI вычисляет индекс массы тела
// weightKg - вес в килограммах
// heightCm - рост в сантиметрах
func CalculateBMI(weightKg, heightCm float64) float64 {
	// TODO: реализуй функцию
	// Не забудь перевести сантиметры в метры!
	// Формула: ИМТ = вес (кг) / рост (м)²
	return 0
}

// InterpretBMI возвращает категорию по значению ИМТ
func InterpretBMI(bmi float64) string {
	// TODO: реализуй функцию
	// < 18.5: "Недостаточный вес"
	// 18.5 - 24.9: "Норма"
	// 25.0 - 29.9: "Избыточный вес"
	// >= 30: "Ожирение"
	return ""
}

func main() {
	// Тест 1: вес 70 кг, рост 175 см → ИМТ ≈ 22.9 (Норма)
	bmi := CalculateBMI(70, 175)
	fmt.Printf("Вес: 70 кг, Рост: 175 см\n")
	fmt.Printf("ИМТ: %.1f\n", bmi)
	fmt.Printf("Категория: %s\n\n", InterpretBMI(bmi))

	// TODO: Добавь ещё несколько тестов:
	// - Недостаточный вес
	// - Избыточный вес
	// - Ожирение
}
