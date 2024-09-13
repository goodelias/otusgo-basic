package shape

import "fmt"

type triangle struct {
	base, height float64
}

func NewTriangle(base, height float64) Shape {
	return &triangle{base: base, height: height}
}

func (t *triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t *triangle) String() string {
	return fmt.Sprintf("Треугольник: основание %.2f, высота %.2f", t.base, t.height)
}
