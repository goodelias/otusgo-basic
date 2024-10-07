package shape

import (
	"errors"
	"fmt"
)

type rectangle struct {
	width  float64
	height float64
}

var errNegativeWidthHeight = errors.New("width and height cannot be negative")

func NewRectangle(width float64, height float64) (Shape, error) {
	if width < 0 || height < 0 {
		return nil, errNegativeWidthHeight
	}
	return &rectangle{width: width, height: height}, nil
}

func (r *rectangle) Area() float64 {
	return r.width * r.height
}

func (r *rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %.2f, высота %.2f", r.width, r.height)
}
