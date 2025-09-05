package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 1. Выход по условию
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("По условию: работа %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("По условию: завершено")
	}()

	// 2. Через канал done
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("По каналу: завершено")
				return
			default:
				fmt.Println("По каналу: работа")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// 3. При закрытии канала
	dataChan := make(chan int)
	go func() {
		for data := range dataChan {
			fmt.Printf("По закрытию: получил %d\n", data)
		}
		fmt.Println("По закрытию: завершено")
	}()

	// 4. Через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("По контексту: завершено")
				return
			default:
				fmt.Println("По контексту: работа")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// 5. Контекст с таймаутом
	timeoutCtx, _ := context.WithTimeout(context.Background(), 800*time.Millisecond)
	go func() {
		for {
			select {
			case <-timeoutCtx.Done():
				fmt.Println("По таймауту: завершено")
				return
			default:
				fmt.Println("По таймауту: работа")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// 6. Через runtime.Goexit()
	go func() {
		defer fmt.Println("По Goexit: defer выполнен")
		fmt.Println("По Goexit: запущено")
		runtime.Goexit()
	}()

	// 7. Через WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("По WaitGroup: работа")
		fmt.Println("По WaitGroup: завершено")
	}()

	// 8. Через panic/recover
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("По panic: восстановлено")
			}
		}()
		panic("аварийная остановка")
	}()

	// Даем время на работу
	time.Sleep(300 * time.Millisecond)
	dataChan <- 1
	dataChan <- 2
	close(dataChan)

	time.Sleep(300 * time.Millisecond)
	close(done)
	cancel()

	time.Sleep(300 * time.Millisecond)
	wg.Wait()

	fmt.Println("\nВсе горутины завершены")
}
