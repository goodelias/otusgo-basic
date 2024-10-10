package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name    string
		arr     []int
		target  int
		want    int
		wantErr error
	}{
		{
			name:    "with existing target",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  3,
			want:    2,
			wantErr: nil,
		},
		{
			name:    "with not existing target",
			arr:     []int{1, 2, 3, 4, 5},
			target:  6,
			want:    -1,
			wantErr: ErrElementNotFound,
		},
		{
			name:    "with empty array",
			arr:     []int{},
			target:  1,
			want:    -1,
			wantErr: ErrElementNotFound,
		},
		{
			name:    "with first target element",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  1,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with last target element",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  9,
			want:    8,
			wantErr: nil,
		},
		{
			name:    "with single target element",
			arr:     []int{5},
			target:  5,
			want:    0,
			wantErr: nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := binarySearch(test.arr, test.target)
			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, -1, got)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
