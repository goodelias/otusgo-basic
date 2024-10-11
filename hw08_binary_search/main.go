package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	index := binarySearch(s, target)
	if index == -1 {
		fmt.Printf("Element %d not found", target)
		return
	}
	fmt.Printf("Element %d found at index %d\n", target, index)
}
