package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChess(t *testing.T) {
	testCases := []struct {
		name    string
		size    int
		want    string
		wantErr error
	}{
		{
			name:    "3x3 board",
			size:    3,
			want:    "# #\n # \n# #\n",
			wantErr: nil,
		},
		{
			name:    "4x4 board",
			size:    4,
			want:    "# # \n # #\n# # \n # #\n",
			wantErr: nil,
		},
		{
			name:    "with too small size",
			size:    -1,
			want:    "",
			wantErr: errInvalidSize,
		},
		{
			name:    "with too large size",
			size:    101,
			want:    "",
			wantErr: errInvalidSize,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := displayBoard(test.size)
			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, "", got)
				assert.EqualError(t, err, test.wantErr.Error())
			}
			assert.Equal(t, test.want, got)
		})
	}
}
