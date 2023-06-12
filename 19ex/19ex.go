package main

import "fmt"

func main() {
	s := "главрыба"

	fmt.Printf("%s", Reverse(s))
}

func Reverse(s string) string {
	c := []rune(s)

	length := len(c)

	for i := 0; i < length/2; i++ {
		c[i], c[length-i-1] = c[length-i-1], c[i]
	}

	return string(c)
}
