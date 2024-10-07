package shape

import (
	"errors"
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

var errNegativeRadius = errors.New("radius cannot be negative")

func NewCircle(radius float64) (Shape, error) {
	if radius < 0 {
		return nil, errNegativeRadius
	}

	return &circle{radius: radius}, nil
}

func (c *circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) String() string {
	return fmt.Sprintf("Круг: радиус %.2f", c.radius)
}
