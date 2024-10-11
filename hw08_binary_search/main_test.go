package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		target int
		want   int
		found  bool
	}{
		{
			name:   "with existing target",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 3,
			want:   2,
			found:  true,
		},
		{
			name:   "with not existing target",
			arr:    []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
			found:  false,
		},
		{
			name:   "with empty array",
			arr:    []int{},
			target: 1,
			want:   -1,
			found:  false,
		},
		{
			name:   "with first target element",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 1,
			want:   0,
			found:  true,
		},
		{
			name:   "with last target element",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 9,
			want:   8,
			found:  true,
		},
		{
			name:   "with single target element",
			arr:    []int{5},
			target: 5,
			want:   0,
			found:  true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, found := binarySearch(test.arr, test.target)
			if test.found != true {
				assert.False(t, found)
				assert.Equal(t, -1, got)
				return
			}
			assert.True(t, found)
			assert.Equal(t, test.want, got)
		})
	}
}
