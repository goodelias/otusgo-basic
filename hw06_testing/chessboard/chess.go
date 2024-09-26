package chessboard

import (
	"errors"
	"fmt"
)

var errInvalidSize = errors.New("size must be between 0 and 100")

func GetSize(prompt string) int {
	var size int
	fmt.Print(prompt)
	for {
		_, err := fmt.Scanf("%d", &size)
		if err == nil && size >= 1 && size <= 100 {
			break
		}
		fmt.Println("Неверный ввод, попробуйте еще раз.")
	}
	return size
}

func displayBoard(size int) (string, error) {
	var board string
	if size < 0 || size > 100 {
		return "", errInvalidSize
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board += "#"
				continue
			}
			board += " "
		}
		board += "\n"
	}
	return board, nil
}
