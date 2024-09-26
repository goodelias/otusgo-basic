package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle_Area(t *testing.T) {
	testCases := []struct {
		name    string
		width   float64
		height  float64
		want    float64
		wantErr error
	}{
		{
			name:    "with normal width and height",
			width:   3.0,
			height:  2.0,
			want:    3.0 * 2.0,
			wantErr: nil,
		},
		{
			name:    "with negative width and height",
			width:   -4,
			height:  -2,
			want:    0,
			wantErr: errNegativeWidthHeight,
		},
		{
			name:    "with zero width and height",
			width:   0,
			height:  0,
			want:    0,
			wantErr: nil,
		},
		{
			name:    "with very large width and height",
			width:   1e20,
			height:  1e20,
			want:    1e20 * 1e20,
			wantErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rectangle, err := NewRectangle(test.width, test.height)

			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Nil(t, rectangle)
				assert.Equal(t, test.wantErr, err)
				return
			}
			got := rectangle.Area()
			assert.Equal(t, test.want, got)
		})
	}
}
