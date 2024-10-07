package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateArea(t *testing.T) {
	circle, err := NewCircle(3)
	assert.NoError(t, err)

	testCases := []struct {
		name    string
		shape   any
		want    float64
		wantErr error
	}{
		{
			name:    "valid shape",
			shape:   circle,
			want:    28.27,
			wantErr: nil,
		},
		{
			name:    "invalid shape",
			shape:   "",
			want:    0,
			wantErr: errInvalidShape,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculateArea(test.shape)

			if test.wantErr != nil {
				assert.Equal(t, test.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.InDelta(t, test.want, got, 0.01)
		})
	}
}
