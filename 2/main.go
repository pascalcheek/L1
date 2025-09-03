package main

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squares := calculateSquares(nums)

	for _, v := range squares {
		fmt.Printf("%d ", v)
	}
}

func calculateSquares(nums []int) []int {
	// создаем waitgroup, чтобы синхронизировать потоки(в противном случае результат мог быть каким угодно
	// мы бы не дожидались выполнения всех горутин и выходили из функции
	wg := sync.WaitGroup{}
	// waitgroup работает как счетчик, если баланс счетчика становится равным 0,
	//то он пропускает наш основной(calculateSquares, он же и в main) поток дальше wg.Wait()
	// тут я бы мог написать wg.Add(len(nums)), но мне больше нравится писать wg.Add(1) в цикле
	for i, num := range nums {
		wg.Add(1)
		go calculateSquare(num, &nums[i], &wg)
	}

	wg.Wait() // ждем исполнения calculateSquare всеми потоками
	return nums
}

func calculateSquare(num int, pos *int, wg *sync.WaitGroup) {
	defer wg.Done() // -1 от счетчика
	*pos = num * num
}
