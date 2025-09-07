package main

import "fmt"

// BinSearch выполняет бинарный поиск в отсортированном срезе.
// Возвращает индекс элемента или -1, если элемент не найден.
func BinSearch(arr []int, x int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := l + (r-l)/2 // Предотвращает переполнение

		if arr[m] == x {
			return m // Элемент найден
		} else if arr[m] < x {
			l = m + 1 // Искомый элемент в правой половине
		} else {
			r = m - 1 // Искомый элемент в левой половине
		}
	}

	return -1 // Элемент не найден
}

func main() {
	// Тестирование алгоритма
	testCases := []struct {
		arr      []int
		x        int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, 3, -1},
		{[]int{1, 3, 5, 7, 9}, 7, 3},
		{[]int{1, 3, 5, 7}, 3, 1},
	}

	for i, tc := range testCases {
		res := BinSearch(tc.arr, tc.x)
		status := "ok"
		if res != tc.expected {
			status = "error"
		}
		fmt.Printf("Тест %d %s: arr=%v, x=%d, expected=%d, got=%d\n",
			i+1, status, tc.arr, tc.x, tc.expected, res)
	}
}
