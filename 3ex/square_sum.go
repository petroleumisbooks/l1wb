package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	var sum = 0

	// создаем мьютекс, чтобы мы могли конкурентно записывать данные в переменную
	mu := sync.Mutex{}

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			sum += val * val
			mu.Unlock()
		}(arr[i])
	}
	wg.Wait()
	fmt.Println(sum)
}
