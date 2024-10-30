package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []float32
	}{
		{
			name:     "10 numbers",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []float32{5.5},
		},
		{
			name:     "20 numbers",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: []float32{5.5, 15.5},
		},
		{
			name:     "0 numbers",
			numbers:  []int{},
			expected: []float32{},
		},
		{
			name:     "3 numbers",
			numbers:  []int{1, 2, 3},
			expected: []float32{2.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int)
			out := make(chan float32)
			go func() {
				for _, num := range tt.numbers {
					in <- num
				}
				close(in)
			}()

			go CalculateAverage(in, out)

			var result []float32
			for avg := range out {
				result = append(result, avg)
			}

			if result == nil {
				result = []float32{}
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
