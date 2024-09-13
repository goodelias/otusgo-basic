package shape

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width float64, height float64) Shape {
	return &rectangle{width: width, height: height}
}

func (r *rectangle) Area() float64 {
	return r.width * r.height
}

func (r *rectangle) String() string {
	return fmt.Sprintf("Прямоугольник: ширина %.2f, высота %.2f", r.width, r.height)
}
