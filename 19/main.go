package main

import (
	"fmt"
)

// Reverse переворачивает строку с учётом Unicode символов
func Reverse(s string) string {
	// Конвертируем строку в срез рун для корректной работы с Unicode
	runes := []rune(s)

	// Переворачиваем срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
		// мне нравится такой способ, тк он не требует выделения доп. памяти
	}

	// Конвертируем обратно в строку
	return string(runes)
}

func main() {
	// Тестовые примеры
	tests := []string{
		"главрыба",
		"",
		"главрыба世界!",
		"a",
		"hello",
		"12345",
	}

	for _, test := range tests {
		reversed := Reverse(test)
		fmt.Printf("Оригинал: \"%s\"\n", test)
		fmt.Printf("Перевернутая: \"%s\"\n", reversed)
		fmt.Println()
	}
}
