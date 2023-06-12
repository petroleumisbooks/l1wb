package main

import (
	"fmt"
	"sync"
	"time"
)

// остановка с помощью закрытия канала

func main() {
	stopCh := make(chan interface{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	go someFunc(stopCh, &wg)

	time.Sleep(5 * time.Second)

	fmt.Println("its TIIIIIIIIIIME")
	close(stopCh)

	wg.Wait()
}

func someFunc(stopCh chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sigma := 6

	for {
		select {
		case <-stopCh:
			return
		default:
			sigma += sigma
			time.Sleep(1 * time.Second)
		}
	}
}
