package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	workersNumber := flag.Int("workers", 10, "Number of goroutines")
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(*workersNumber)

	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, os.Interrupt)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-chExit:
				close(ch)
				return
			default:
				ch <- rand.Intn(5000)
			}
		}
	}()

	for i := 0; i < *workersNumber; i++ {
		go func(num int) {
			defer wg.Done()
			for {
				if val, ok := <-ch; ok {
					//worker number and the value
					fmt.Printf(" %d : %d\n", num, val)
					continue
				}
				return
			}
		}(i + 1)
	}

	wg.Wait()
}
