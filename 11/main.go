package main

import "fmt"

// Т.к. в условии ничего не сказано про числа, то я решил сделать дженериками(comparable)
// Есть несколько вариантов решения:
// 1) Отсортировать оба списка и пройтись двумя указателями (время: O(nlogn), память: O(1))
// 2) То что я реализовал. Записать значения одного списка как ключ map, а потом пройтись по второму и проверить наличие
// (время: O(n), память: O(n))

func main() {
	// Пример с числами
	A1 := []int{1, 2, 3, 4, 5, 2, 3}
	B1 := []int{4, 5, 6, 7, 8, 4, 5}
	fmt.Printf("Числа: %v ∩ %v = %v\n", A1, B1, Intersect(A1, B1))

	// Пример со строками
	A2 := []string{"apple", "banana", "orange", "banana"}
	B2 := []string{"banana", "kiwi", "orange", "kiwi"}
	fmt.Printf("Строки: %v ∩ %v = %v\n", A2, B2, Intersect(A2, B2))

	// Пример с float
	A3 := []float64{1.1, 2.2, 3.3, 4.4}
	B3 := []float64{2.2, 3.3, 5.5, 6.6}
	fmt.Printf("Float: %v ∩ %v = %v\n", A3, B3, Intersect(A3, B3))
}

// Intersect возвращает пересечение двух множеств для любого comparable типа
func Intersect[T comparable](a, b []T) []T {
	// Создаем множество из первого слайса
	set := make(map[T]bool)
	for _, x := range a {
		set[x] = true
	}

	// Ищем общие элементы во втором слайсе
	var res []T
	for _, x := range b {
		if set[x] {
			res = append(res, x)
			set[x] = false // Помечаем как уже добавленное
		}
	}

	return res
}
