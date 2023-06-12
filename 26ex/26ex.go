package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("unique (\"abcd\") or not? : %v ", unique("abcd"))
}

func unique(s string) bool {
	symbolsMap := make(map[rune]struct{})

	for _, r := range s {
		lowRune := unicode.ToLower(r)
		if _, ok := symbolsMap[lowRune]; ok {
			return false
		}
		symbolsMap[lowRune] = struct{}{}
	}

	return true
}
