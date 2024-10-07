package shape

import (
	"errors"
	"fmt"
)

type triangle struct {
	base, height float64
}

var errNegativeBaseHeight = errors.New("base and height cannot be negative")

func NewTriangle(base, height float64) (Shape, error) {
	if base < 0 || height < 0 {
		return nil, errNegativeBaseHeight
	}
	return &triangle{base: base, height: height}, nil
}

func (t *triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t *triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %.2f, высота %.2f", t.base, t.height)
}
