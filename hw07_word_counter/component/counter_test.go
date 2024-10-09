package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testText = `In a bustling city, there lived a young artist named Jane, 
passionate about painting. One day, while wandering vibrant streets, 
she discovered a hidden alley filled with colorful murals.
Inspired, she set up her easel and began to paint. Curious onlookers gathered,
admiring her talent. As the sun set, casting a golden hue, 
Jane stepped back to admire her creation, reflecting on her journey. 
"Life is like a canvas," she thought, 
"every moment is an opportunity to create something beautiful."`

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name string
		text string
		want map[string]int
	}{
		{
			name: "with many repeated words",
			text: "Go go go! It's time to go home. Are you ready to go? Let's go!",
			want: map[string]int{
				"go":    6,
				"to":    2,
				"it's":  1,
				"time":  1,
				"home":  1,
				"are":   1,
				"you":   1,
				"ready": 1,
				"let's": 1,
			},
		},
		{
			name: "with other language",
			text: "Что ты думаешь об ИИ? Это наше будущее? Я отвечу - да! А что думаешь ты?",
			want: map[string]int{
				"что":     2,
				"ты":      2,
				"думаешь": 2,
				"об":      1,
				"ии":      1,
				"это":     1,
				"наше":    1,
				"будущее": 1,
				"я":       1,
				"отвечу":  1,
				"да":      1,
				"а":       1,
			},
		},
		{
			name: "with empty text",
			text: "",
			want: map[string]int{},
		},
		{
			name: "with numbers in words",
			text: "hi bro007, how are you b2ro?",
			want: map[string]int{"bro007": 1, "b2ro": 1, "hi": 1, "how": 1, "are": 1, "you": 1},
		},
		{
			name: "with special symbols",
			text: "Hey! Did you see the new @Symbol#? It costs $100, but they say it's worth it.",
			want: map[string]int{
				"it":     2,
				"hey":    1,
				"did":    1,
				"you":    1,
				"see":    1,
				"the":    1,
				"new":    1,
				"symbol": 1,
				"costs":  1,
				"100":    1,
				"but":    1,
				"they":   1,
				"say":    1,
				"it's":   1,
				"worth":  1,
			},
		},
		{
			name: "with very long text",
			text: testText,
			want: map[string]int{
				"a":           5,
				"her":         4,
				"she":         3,
				"to":          3,
				"jane":        2,
				"set":         2,
				"is":          2,
				"in":          1,
				"bustling":    1,
				"city":        1,
				"there":       1,
				"lived":       1,
				"young":       1,
				"artist":      1,
				"named":       1,
				"passionate":  1,
				"about":       1,
				"painting":    1,
				"one":         1,
				"day":         1,
				"while":       1,
				"wandering":   1,
				"vibrant":     1,
				"streets":     1,
				"discovered":  1,
				"hidden":      1,
				"alley":       1,
				"filled":      1,
				"with":        1,
				"colorful":    1,
				"murals":      1,
				"inspired":    1,
				"up":          1,
				"easel":       1,
				"and":         1,
				"began":       1,
				"paint":       1,
				"curious":     1,
				"onlookers":   1,
				"gathered":    1,
				"admiring":    1,
				"talent":      1,
				"as":          1,
				"the":         1,
				"sun":         1,
				"casting":     1,
				"golden":      1,
				"hue":         1,
				"stepped":     1,
				"back":        1,
				"admire":      1,
				"creation":    1,
				"reflecting":  1,
				"on":          1,
				"journey":     1,
				"life":        1,
				"like":        1,
				"canvas":      1,
				"thought":     1,
				"every":       1,
				"moment":      1,
				"an":          1,
				"opportunity": 1,
				"create":      1,
				"something":   1,
				"beautiful":   1,
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := CountWords(test.text)
			assert.Equal(t, test.want, got)
		})
	}
}
