package main

import (
	"errors"
	"fmt"

	"github.com/goodelias/otusgo-basic/hw05_shapes/shape"
)

func calculateArea(s any) (float64, error) {
	figure, ok := s.(shape.Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return figure.Area(), nil
}

func main() {
	circle := shape.NewCircle(5)
	rectangle := shape.NewRectangle(3, 5)
	triangle := shape.NewTriangle(7, 3)

	figures := []any{circle, rectangle, triangle, 4}

	for _, figure := range figures {
		area, err := calculateArea(figure)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		fmt.Printf("%s Площадь: %.2f\n", figure, area)
	}
}
