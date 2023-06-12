package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 12, 253}
	element := 12
	BinarySearch(arr, element)

}

func BinarySearch(arr []int, el int) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if el < arr[mid] {
			right = mid - 1
		} else if el > arr[mid] {
			left = mid + 1
		} else {
			fmt.Println(mid)
			return
		}
	}
	fmt.Print(-1)
}
