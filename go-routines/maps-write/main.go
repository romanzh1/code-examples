package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000

	mu := sync.RWMutex{}
	// ch := make(chan struct{}, 1)

	storage := make(map[int]int, 0)
	wg := sync.WaitGroup{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			// ch <- struct{}{}
			storage[i] = i
			// <-ch
			mu.Unlock()
		}(i)
	}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		go func(i int) {
			defer wg.Done()

			mu.RLock()
			// ch <- struct{}{}
			fmt.Println(storage[i])
			// <-ch
			mu.RUnlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("OK")
}
