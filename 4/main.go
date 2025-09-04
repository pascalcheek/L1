package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Я выбрал способ с Context, тк он является стандартным.
// Его используют в проектах, тк есть иерархическое отменение и это упрощает управление

func main() {
	// запрашиваем количество воркеров через консоль
	workersSize := getWorkerCount()

	// создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Гарантируем освобождение ресурсов

	ch := make(chan string)
	var wg sync.WaitGroup

	// запускаем воркеры
	for i := 0; i < workersSize; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	// обработка сигналов для корректного завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// главная горутина записывает данные в канал
	go func() {
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				message := fmt.Sprintf("Message %d", counter)
				// правильно писать этот select. Вот почему:
				// представим себе что мы прочитали из message какое-то значение
				// у нас в это время ctx.Done() какое-то значение вернул, и мы в ch пытаемся что-то записать
				// Но может быть такое что цепочка событий которая происходит дальше может не читать ничего из ch и
				// и мы бы заблокировались на строке ch <- message
				// Так что принято писать еще вот этот вложенный select. Это стандартный паттерн
				select {
				case ch <- message:
					counter++
				case <-ctx.Done():
					return
				}

				// Эта строка для того чтобы было легче наблюдать за работой
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// ожидаем сигнал завершения
	<-sigCh

	// отменяем контекст - сигнализируем всем горутинам о завершении
	cancel()

	// закрываем канал после отмены контекста
	close(ch)

	// ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers done")
}

// worker - функция воркера с использованием контекста
func worker(ctx context.Context, id int, dataChannel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)

	for {
		select {
		case <-ctx.Done():
			// получен сигнал отмены контекста
			fmt.Printf("Worker %d: closing\n", id)
			return

		case message, ok := <-dataChannel:
			if !ok {
				// канал закрыт
				fmt.Printf("Worker %d: done\n", id)
				return
			}
			fmt.Printf("Worker %d: %s\n", id, message)
		}
	}
}

// getWorkerCount запрашивает у пользователя количество воркеров
func getWorkerCount() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Write workers count: ")
		input, _ := reader.ReadString('\n')

		// убираем символы новой строки
		if len(input) > 0 {
			if input[len(input)-1] == '\n' {
				input = input[:len(input)-1]
			}
			if len(input) > 0 && input[len(input)-1] == '\r' {
				input = input[:len(input)-1]
			}
		}

		if input == "" {
			fmt.Println("Write a number")
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Write a number")
			continue
		}

		if num <= 0 {
			fmt.Println("Workers count must be positive")
			continue
		}

		return num
	}
}
