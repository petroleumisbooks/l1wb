package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	mainChan := make(chan int, 10)
	squareChan := make(chan int, 10)

	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, os.Interrupt)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-chExit:
				close(mainChan)
				return
			default:
				mainChan <- rand.Intn(500)
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for x := range mainChan {
			squareChan <- x * x
		}
		close(squareChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range squareChan {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
