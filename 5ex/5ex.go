package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	seconds := flag.Int("seconds", 3, "How much programm should work")
	flag.Parse()

	ch := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*seconds)*time.Second)
	defer cancel()

	// далее мы создаем две горутины:
	// одна записывает в канал, вторая считывает

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- rand.Intn(5000)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if val, ok := <-ch; ok {
				fmt.Println(val)
			} else {
				return
			}
		}
	}()

	wg.Wait()
}
