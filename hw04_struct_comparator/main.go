package main

import (
	"fmt"
	"github.com/goodelias/otusgo-basic/hw04_struct_comparator/component"
	"github.com/goodelias/otusgo-basic/hw04_struct_comparator/entity"
)

func main() {
	book1 := entity.NewBookBuilder().
		SetRate(4.5).
		SetId(1).
		SetAuthor("Ilya Chepkin").
		SetSize(300).
		SetTitle("Go Basic").
		SetYear(2015).
		Build()

	book2 := entity.NewBookBuilder().
		SetId(2).
		SetTitle("Go Professional").
		SetAuthor("Ilya Chepkin").
		SetYear(2017).
		SetSize(400).
		SetRate(4.7).
		Build()

	bc1 := component.NewBookComparator(component.CompareModeRate)
	fmt.Println(bc1.Compare(book1, book2))
	book1.SetRate(4.8)
	fmt.Println(bc1.Compare(book1, book2))

}
