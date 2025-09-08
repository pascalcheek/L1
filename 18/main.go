package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter счётчик на atomic операциях
type Counter struct {
	value int64
}

func NewCounter() *Counter {
	return &Counter{0} // Вообще 0 можно было не устанавливать, это значение по умолчанию
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Counter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	// Тест счётчика
	counter := NewCounter()
	var wg sync.WaitGroup

	// Запускаем 100 горутин
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Итоговое значение: %d\n", counter.GetValue())
	fmt.Printf("Ожидаемое значение: %d\n", 100*100)
}
