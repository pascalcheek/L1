package main

import (
	"fmt"
)

func main() {
	var num int64
	var pos uint
	var bitFlag int
	// Обработки неверных данных нет. Думаю задание не про это
	// Ввод данных
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита (0-63): ")
	fmt.Scan(&pos)

	fmt.Print("Установить бит в (0 или 1): ")
	fmt.Scan(&bitFlag)

	// Установка бита
	result := SetBit(num, pos, bitFlag)

	// Вывод результатов
	fmt.Printf("\nИсходное: %d (%064b)\n", num, num)
	fmt.Printf("Установка %d-го бита в %d\n", pos, bitFlag)
	fmt.Printf("Результат: %d (%064b)\n", result, result)
}

func SetBit(num int64, pos uint, bitFlag int) int64 {
	// Создаем маску: 1 сдвинутый на i позиций влево
	// Например, для i=5: 1 << 5 = 32 (100000 в двоичной)
	mask := int64(1 << pos)

	if bitFlag == 1 {
		// Установка бита в 1: используем OR
		// OR оставляет все биты как есть, но устанавливает i-й в 1
		return num | mask
	} else {
		// Установка бита в 0: используем AND NOT
		// AND NOT оставляет все биты как есть, но устанавливает i-й в 0
		return num &^ mask
	}
}
