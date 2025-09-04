package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// Количество секунд, через которое программа завершается
	timeoutInSecs := getTimeout()

	// Канал для передачи значений
	ch := make(chan any)

	// Запуск писателя с таймаутом
	go produce(ch, timeoutInSecs)

	// Создание читателя значения из канала ch
	consume(ch)
}

func getTimeout() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Write N: ")
		input, _ := reader.ReadString('\n')

		// Убираем символы новой строки
		if len(input) > 0 {
			if input[len(input)-1] == '\n' {
				input = input[:len(input)-1]
			}
			if len(input) > 0 && input[len(input)-1] == '\r' {
				input = input[:len(input)-1]
			}
		}

		if input == "" {
			fmt.Println("Write number")
			continue
		}

		seconds, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Write number")
			continue
		}

		if seconds <= 0 {
			fmt.Println("Write positive number")
			continue
		}

		return seconds
	}
}

// Писатель в канал ch
func produce(ch chan<- any, timeoutInSecs int) {
	timeout := time.After(time.Duration(timeoutInSecs) * time.Second)

	for {
		select {
		case <-timeout: // Проверка времени
			close(ch)
			return
		default:
			select {
			case ch <- generateValue():
				// Успешная запись
			case <-timeout: // Защита от блокировки. Если вдруг больше не читают из канала
				close(ch)
				return
			}
		}

	}
}

// Функция для создания значения для канала
func generateValue() interface{} {
	return struct{ x string }{x: time.Now().Format("2006-01-02 15:04:05")}
}

// Читатель из канал
func consume(ch <-chan any) {
	for message := range ch {
		fmt.Println(message)
	}
}
