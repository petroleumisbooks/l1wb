package main

import "fmt"

func main() {
	a, b := 15, 30
	swap(&a, &b)
	fmt.Print(a, b)
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
