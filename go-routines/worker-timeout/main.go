package main

import (
	"fmt"
	"time"
)

func worker(work <-chan int) {
	fmt.Println(1)
	//count := 0
	for {
		timeout := time.After(4 * time.Second)
		fmt.Println(2)

		select {
		case res := <-work:
			time.Sleep(5 * time.Second)
			fmt.Println(res)
		// ... выполняем что-нибудь

		case <-timeout:
			// Закрываем эту горутину после указанного таймаута.
			fmt.Println("close")
			return
			//default: // select будет ждать, а вот с default он будет бесконечно крутиться
			//	count++
		}
		//fmt.Println(count)
	}
}

func main() {
	work := make(chan int)
	go worker(work)

	time.Sleep(1 * time.Second)
	work <- 3
	fmt.Println("ready")
	time.Sleep(10 * time.Second)
}
