package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// TODO make rw mutex work

// selected only unique values
// find bug with reads-writes
func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.RWMutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(13)) // create rand num 0...9
	}
	// 3, 4, 5, 0, 4, 9, 8, 6, 6, 5, 5, 4, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			// mu.RLock()
			// defer func() {
			// 	if !mu.TryRLock() {
			// 		mu.Unlock()
			// 	}
			// }()

			if _, ok := alreadyStored[doubles[i]]; !ok {
				// mu.RUnlock()

				mu.Lock()
				defer mu.Unlock()

				alreadyStored[doubles[i]] = struct{}{}

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		for val := range uniqueIDs {
			fmt.Println(val)
		}

		wg2.Done()
	}()

	wg.Wait()
	close(uniqueIDs)

	wg2.Wait()

	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0, 1, 2, 3, 4 ...
	fmt.Println(uniqueIDs)
}
