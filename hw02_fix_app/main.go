package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

var ErrBadInput = errors.New("unexpected newline")

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")

	if _, err := fmt.Scanln(&path); err != nil {
		if err.Error() != "unexpected newline" {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printer.PrintStaff(staff)
}
