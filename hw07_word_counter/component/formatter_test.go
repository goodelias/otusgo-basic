package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintFormattedWordCount(t *testing.T) {
	testCases := []struct {
		name         string
		wordCountMap map[string]int
		want         string
		wantErr      error
	}{
		{
			name:         "with basic map of words",
			wordCountMap: map[string]int{"hello": 2, "world": 1, "go": 3},
			want:         "Word: go      | Count: 3\nWord: hello   | Count: 2\nWord: world   | Count: 1\n",
			wantErr:      nil,
		},
		{
			name:         "with empty map",
			wordCountMap: map[string]int{},
			want:         "",
			wantErr:      ErrEmptyMap,
		},
		{
			name:         "with zero count",
			wordCountMap: map[string]int{"hello": 2, "test": 0},
			want:         "Word: hello   | Count: 2\n",
			wantErr:      nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := PrintFormattedWordCount(test.wordCountMap)

			if test.wantErr != nil {
				assert.Error(t, err)
				assert.Empty(t, got)
				assert.EqualError(t, err, ErrEmptyMap.Error())
			}
			assert.Equal(t, test.want, got)
		})
	}
}
