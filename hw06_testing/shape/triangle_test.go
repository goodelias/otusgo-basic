package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangle_Area(t *testing.T) {
	testCases := []struct {
		name    string
		base    float64
		height  float64
		want    float64
		wantErr error
	}{
		{
			name:    "with normal base and height",
			base:    3,
			height:  5,
			want:    0.5 * 3 * 5,
			wantErr: nil,
		},
		{
			name:    "with negative base and height",
			base:    -2,
			height:  -4,
			want:    0,
			wantErr: errNegativeBaseHeight,
		},
		{
			name:    "with zero base and height",
			base:    0,
			height:  0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with very small base and height",
			base:    0.000003,
			height:  0.0000001,
			want:    0.5 * 0.000003 * 0.0000001,
			wantErr: nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			triangle, err := NewTriangle(test.base, test.height)
			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Nil(t, triangle)
				assert.Equal(t, test.wantErr, err)
				return
			}
			got := triangle.Area()
			assert.Equal(t, got, test.want)
		})
	}
}
