package entity

import "errors"

type CompareMode int

const (
	CompareModeYear CompareMode = iota
	CompareModeSize
	CompareModeRate
)

type BookComparator struct {
	mode CompareMode
}

var errEmptyBook = errors.New("one or both books are nil")

func NewBookComparator(mode CompareMode) *BookComparator {
	return &BookComparator{mode}
}

func (c BookComparator) Compare(book1, book2 *Book) (bool, error) {
	if *book1 == (Book{}) && *book2 == (Book{}) {
		return false, errEmptyBook
	}
	switch c.mode {
	case CompareModeYear:
		return book1.Year() > book2.Year(), nil
	case CompareModeSize:
		return book1.Size() > book2.Size(), nil
	case CompareModeRate:
		return book1.Rate() > book2.Rate(), nil
	default:
		return book1.Year() > book2.Year(), nil
	}
}
