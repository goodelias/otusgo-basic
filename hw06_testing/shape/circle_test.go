package shape

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle_Area(t *testing.T) {
	testCases := []struct {
		name    string
		radius  float64
		want    float64
		wantErr error
	}{
		{
			name:    "with normal radius",
			radius:  5,
			want:    math.Pi * 5 * 5,
			wantErr: nil,
		},
		{
			name:    "with zero radius",
			radius:  0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with negative radius",
			radius:  -3,
			want:    0,
			wantErr: errNegativeRadius,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			circle, err := NewCircle(test.radius)

			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Nil(t, circle)
				assert.Equal(t, test.wantErr, err)
				return
			}
			got := circle.Area()
			assert.Equal(t, test.want, got)
		})
	}
}
