package main

import (
	"fmt"

	"github.com/goodelias/otusgo-basic/hw07_word_counter/component"
)

func main() {
	const Text = "Go go go! It's time to go home. Are you ready to go? Let's go!"
	wordsCountMap := component.CountWords(Text)
	formated, err := component.PrintFormattedWordCount(wordsCountMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(formated)
}
