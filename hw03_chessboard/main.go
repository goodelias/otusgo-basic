package main

import "fmt"

func main() {
	var size1 int
	var size2 int

	fmt.Print("Enter size 1 for chessboard\n")
	fmt.Scanf("%d", &size1)
	fmt.Print("Enter size 2 for chessboard\n")
	fmt.Scanf("%d", &size2)

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
