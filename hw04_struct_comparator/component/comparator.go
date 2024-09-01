package component

import "github.com/goodelias/otusgo-basic/hw04_struct_comparator/entity"

type CompareMode int

const (
	CompareModeYear CompareMode = iota
	CompareModeSize
	CompareModeRate
)

type BookComparator struct {
	mode CompareMode
}

func NewBookComparator(mode CompareMode) *BookComparator {
	return &BookComparator{mode}
}

func (c BookComparator) Compare(book1, book2 *entity.Book) bool {
	switch c.mode {
	case CompareModeYear:
		return book1.Year() > book2.Year()
	case CompareModeSize:
		return book1.Size() > book2.Size()
	case CompareModeRate:
		return book1.Rate() > book2.Rate()
	default:
		return book1.Year() > book2.Year()
	}
}
