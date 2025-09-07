package main

import "fmt"

func main() {
	// Исходная последовательность строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество (map для хранения уникальных элементов)
	// В Go нет set, так что используют map
	set := make(map[string]bool)

	// Добавляем все элементы в множество
	for _, item := range arr {
		set[item] = true
	}

	// Получаем уникальные элементы
	var res []string
	for item := range set {
		res = append(res, item)
	}

	// Выводим результат
	fmt.Printf("Исходная последовательность: %v\n", arr)
	fmt.Printf("Собственное множество: %v\n", res)
}
