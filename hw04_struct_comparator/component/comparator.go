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

func (b BookComparator) Compare(book1, book2 *entity.Book) bool {
	switch b.mode {
	case CompareModeYear:
		return book1.GetYear() > book2.GetYear()
	case CompareModeSize:
		return book1.GetSize() > book2.GetSize()
	case CompareModeRate:
		return book1.GetRate() > book2.GetRate()
	default:
		return book1.GetYear() > book2.GetYear()
	}
}
