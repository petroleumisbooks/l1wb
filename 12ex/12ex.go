package main

import "fmt"

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("%v -> %v", a, createSet(a))
}

func createSet(a []string) []string {
	set := make([]string, 0)
	aMap := make(map[string]struct{})

	for _, val := range a {
		aMap[val] = struct{}{}
	}

	for key := range aMap {
		set = append(set, key)
	}

	return set
}
