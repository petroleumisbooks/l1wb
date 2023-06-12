package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(5)

	for _, val := range arr {
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}

	wg.Wait()
}
