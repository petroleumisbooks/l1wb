package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// via context.WithCancel

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println(rand.Intn(5000))
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}
