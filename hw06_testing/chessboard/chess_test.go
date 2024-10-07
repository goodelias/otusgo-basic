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
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := displayBoard(test.size)
			assert.Equal(t, test.want, got)
		})
	}
}
