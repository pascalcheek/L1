package main

import (
	"fmt"
	"strings"
)

// Проблемы и что происходит:
// 1)
// v занимает 1024 байта в памяти
// justString = v[:100] создает новую строку, но
// под капотом строка justString содержит:
// Указатель: на начало массива v
// Длина: 100
// Емкость: 1024 (неявно, через ссылку на исходный массив)
// Результат: весь массив из 1024 байт НЕ может быть освобожден GC,
// потому что justString продолжает на него ссылаться!

// 2) Unicode символы(кириллица, эмодзи, китайский язык и тд) будут обрезаться. Лучше переводить в руны.

// createHugeString создает большую строку
func createHugeString(size int) string {
	return strings.Repeat("世", size) // Unicode
}

// SafeSubstring безопасно извлекает подстроку
func SafeSubstring(s string, length int) string {
	// Для многобайтовых символов используем руны
	runes := []rune(s)

	// Проверяем, чтобы не выйти за границы
	if length > len(runes) {
		length = len(runes)
	}

	// Возвращаем новую строку - создается новый массив в памяти
	return string(runes[:length])
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return SafeSubstring(v, 100) // Ровно 100 символов
}

func main() {
	justString := someFunc()

	fmt.Printf("Результат: %s\n", justString)
	fmt.Printf("Длина в символах: %d\n", len([]rune(justString)))
	fmt.Printf("Длина в байтах: %d\n", len(justString))

	// Демонстрация: каждый символ "世" занимает 3 байта
	// 100 символов × 3 байта = 300 байт
}
