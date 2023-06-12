package main

import (
	"math/rand"
)

var justString string

func main() {
	//someFunc()
	someFunc2()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func createHugeString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func someFunc2() {
	v := createHugeString(1 << 10)
	temp := make([]rune, 100)
	copy(temp, []rune(v[:100]))
	justString = string(temp)
}
