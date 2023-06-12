package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	workers := flag.Int("workers", 10, "How much goroutines will work")

	m := make(map[string]int)
	mu := sync.Mutex{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(*workers)

	for i := 0; i < *workers; i++ {
		go func(num int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					mu.Lock()
					m[fmt.Sprintf("Goroutine number %d", num)]++
					mu.Unlock()
				}
			}
		}(i + 1)
	}
	wg.Wait()
	for key, val := range m {
		fmt.Printf("%s: %d\n", key, val)
	}
}
