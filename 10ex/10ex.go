package main

import "fmt"

func main() {
	arr := []float64{
		-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.5,
	}

	m := subsetCreater(arr)

	fmt.Println(m)
}

func subsetCreater(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)

	for _, val := range arr {
		if val < 0 && int(val/10) == 0 {
			key := -1
			m[key] = append(m[key], val)
			continue
		}
		key := int(val/10) * 10
		m[key] = append(m[key], val)
	}

	return m
}
