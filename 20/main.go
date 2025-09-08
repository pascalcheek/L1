package main

import (
	"fmt"
)

// ReverseWords переворачивает порядок слов в строке на месте
// Мы сначала перевернем всю строку, а потом каждое слово по отдельности
// Функция реверс перворачивает двумя указателями, так что в решении мы избегаем выделения дополнительной памяти
func ReverseWords(s string) string {
	// Конвертируем строку в срез рун для работы на месте
	runes := []rune(s)

	// Сначала переворачиваем всю строку
	reverse(runes, 0, len(runes)-1)

	// Затем переворачиваем каждое слово обратно
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' { // наткнулись на конец слова
			// Переворачиваем слово от start до i-1
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

// reverse переворачивает подстроку в срезе рун от start до end
func reverse(runes []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func main() {
	// Тестовые примеры
	tests := []string{
		"",
		"Hello World!",
		"1 2 3",
		"a b c d e",
		"1",
		"   1      2   ",
		"世界",
		" ",
	}

	for _, test := range tests {
		result := ReverseWords(test)
		fmt.Printf("Изначальная: \"%s\"\n", test)
		fmt.Printf("Финальная:   \"%s\"\n", result)
		fmt.Println()
	}
}
