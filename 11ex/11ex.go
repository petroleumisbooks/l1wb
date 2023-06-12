package main

import "fmt"

func main() {
	a := []int{1, 3, -7, 4}
	b := []int{-13, -32, 0, 3}

	c := intersect(a, b)

	for _, v := range c {
		fmt.Printf("%d ", v)
	}

}

func intersect(a []int, b []int) []int {
	c := make([]int, 0)

	cMap := make(map[int]struct{})

	for _, val := range c {
		cMap[val] = struct{}{}
	}

	for _, val := range b {
		if _, ok := cMap[val]; ok {
			c = append(c, val)
		}
	}

	return c
}
