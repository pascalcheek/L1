package main

import (
	"fmt"
	"sync"
)

// Мы должны делать блокировки, чтобы не допускать состояния гонки

type Map struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewMap() *Map {
	return &Map{data: make(map[string]int)}
}

func (sm *Map) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *Map) Get(key string) (int, bool) {
	sm.mu.RLock() // Блокировка только для чтения
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

func (sm *Map) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

func main() {
	var wg sync.WaitGroup
	m := NewMap()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key_%d", j)
				m.Set(key, id*1000+j)
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Записано элементов: %d\n", m.Len())
}

// pascal@Pascal:/mnt/c/Users/paska/GolandProjects/L1/7$ go run -race main.go
// Записано элементов: 100
