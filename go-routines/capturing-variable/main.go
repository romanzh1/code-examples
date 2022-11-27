package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(i)
		}()
	}

	wg.Wait()
}
