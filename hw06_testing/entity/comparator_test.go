package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var book1 = NewBookBuilder().
	SetID(1).
	SetTitle("Go Basic").
	SetAuthor("Ilya Chepkin").
	SetYear(2015).
	SetSize(450).
	SetRate(4.7).
	Build()

var book2 = NewBookBuilder().
	SetID(2).
	SetTitle("Go Professional").
	SetAuthor("Ilya Chepkin").
	SetYear(2017).
	SetSize(400).
	SetRate(4.7).
	Build()

func TestNewBookComparator_Compare(t *testing.T) {
	testCases := []struct {
		name    string
		b1      *Book
		b2      *Book
		bc      BookComparator
		want    bool
		wantErr error
	}{
		{
			name:    "by Year with normal books",
			b1:      book1,
			b2:      book2,
			bc:      BookComparator{CompareModeYear},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "by Size with normal books",
			b1:      book1,
			b2:      book2,
			bc:      BookComparator{CompareModeSize},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "by Rate with normal books",
			b1:      book1,
			b2:      book2,
			bc:      BookComparator{CompareModeRate},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "with empty 1st book",
			b1:      &Book{},
			b2:      book2,
			bc:      BookComparator{CompareModeYear},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "with empty 2nd book",
			b1:      book1,
			b2:      &Book{},
			bc:      BookComparator{CompareModeRate},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "with both empty books",
			b1:      &Book{},
			b2:      &Book{},
			bc:      BookComparator{CompareModeSize},
			want:    false,
			wantErr: errEmptyBook,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.bc.Compare(test.b1, test.b2)
			if test.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.wantErr.Error())
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
