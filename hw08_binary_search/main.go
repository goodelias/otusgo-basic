package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) (int, bool) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case arr[mid] == target:
			return mid, true
		case arr[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1, false
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9
	index, found := binarySearch(s, target)
	if !found {
		fmt.Printf("Element %d not found", target)
		return
	}
	fmt.Printf("Element %d found at index %d\n", target, index)
}
