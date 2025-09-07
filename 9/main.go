package main

import "fmt"

func main() {
	// Исходный массив чисел
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Создаем каналы
	inCh := make(chan int)
	outCh := make(chan int)

	// Запускаем генератор чисел
	go func() {
		for num := range nums {
			inCh <- num // Отправляем числа в первый канал
		}
		close(inCh) // Закрываем канал после отправки всех чисел
	}()

	// Запускаем обработчик
	go func() {
		for num := range inCh { // Читаем из первого канала
			outCh <- num * 2 // Отправляем результат во второй канал
		}
		close(outCh) // Закрываем второй канал после обработки
	}()

	// Выводим результаты из второго канала
	for res := range outCh {
		fmt.Println(res)
	}

}
