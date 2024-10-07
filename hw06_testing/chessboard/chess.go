package chessboard

import (
	"fmt"
	"strings"
)

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

func displayBoard(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				s.WriteString("#")
				continue
			}
			s.WriteString(" ")
		}
		s.WriteString("\n")
	}
	return s.String()
}
