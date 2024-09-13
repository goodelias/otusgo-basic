package shape

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func NewCircle(radius float64) Shape {
	return &circle{radius: radius}
}

func (c *circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) String() string {
	return fmt.Sprintf("Круг: радиус %.2f", c.radius)
}
