package shape

import (
	"errors"
)

var errInvalidShape = errors.New("переданный объект не является фигурой")

func calculateArea(s any) (float64, error) {
	figure, ok := s.(Shape)
	if !ok {
		return 0, errInvalidShape
	}
	return figure.Area(), nil
}
