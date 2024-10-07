package component

import (
	"regexp"
	"strings"
)

func CountWords(text string) map[string]int {
	wordCountMap := make(map[string]int)

	toLowerText := strings.ToLower(text)
	re := regexp.MustCompile(`\p{L}+'?\p{L}*\d*\p{L}*|\d+\p{L}*`)
	words := re.FindAllString(toLowerText, -1)

	for _, word := range words {
		wordCountMap[word]++
	}
	return wordCountMap
}
