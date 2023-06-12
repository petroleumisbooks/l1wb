package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"

	fmt.Printf("%s", Reverse(s))
}

func Reverse(s string) string {
	arr := strings.Split(s, " ")
	length := len(arr)

	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}

	return strings.Join(arr, " ")
}
