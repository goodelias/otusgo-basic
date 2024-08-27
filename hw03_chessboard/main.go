package main

import "fmt"

func main() {
	var size1 int
	var size2 int

	size1 = getSize("Enter size1 for chessboard from 1 to 100: ")
	size2 = getSize("Enter size2 for chessboard from 1 to 100: ")

	displayBoard(size1, size2)
}

func getSize(prompt string) int {
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

func displayBoard(size1, size2 int) {
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("#")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
