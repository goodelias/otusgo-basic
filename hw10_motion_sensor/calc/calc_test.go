package calc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNumbers(t *testing.T) {
	numChannel := make(chan int)

	go GenerateNumbers(numChannel)

	numCount := 0
	timer := time.NewTimer(5 * time.Second)

	for {
		select {
		case _, ok := <-numChannel:
			if ok {
				numCount++
			} else {
				goto Done
			}
		case <-timer.C:
			goto Done
		}
	}

Done:
	expected := 1
	assert.GreaterOrEqual(t, numCount, expected)
}
