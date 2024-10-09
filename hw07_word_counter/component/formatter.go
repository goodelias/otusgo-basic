package component

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var ErrEmptyMap = errors.New("empty map provided")

type WordCount struct {
	Word  string
	Count int
}

func PrintFormattedWordCount(wordCountMap map[string]int) (string, error) {
	if len(wordCountMap) == 0 {
		return "", ErrEmptyMap
	}

	var wordCounts []WordCount

	for word, count := range wordCountMap {
		if count > 0 {
			wordCounts = append(wordCounts, WordCount{
				Word:  word,
				Count: count,
			})
		}
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	var builder strings.Builder

	for _, wc := range wordCounts {
		if wc.Count == 0 {
			continue
		}
		builder.WriteString(fmt.Sprintf("Word: %-7s | Count: %d\n", wc.Word, wc.Count))
	}

	return builder.String(), nil
}
