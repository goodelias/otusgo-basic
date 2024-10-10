package main

import (
	"errors"
	"fmt"
)

var ErrElementNotFound = errors.New("element not found")

func binarySearch(arr []int, target int) (int, error) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case arr[mid] == target:
			return mid, nil
		case arr[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1, ErrElementNotFound
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	search, err := binarySearch(s, target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Element %d found at index %d\n", target, search)
}
