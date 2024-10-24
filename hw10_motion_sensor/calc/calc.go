package calc

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNumbers(ch chan<- int) {
	timer := time.NewTimer(time.Minute)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			num := rand.Intn(10) //#nosec
			fmt.Printf("Generated new number: %d\n", num)
			ch <- num
		case <-timer.C:
			close(ch)
			fmt.Println("Time is up. Generation stopped")
			return
		}
	}
}

func CalculateAverage(numChan <-chan int, avrChan chan<- float32) {
	numbers := make([]int, 0, 10)
	count := 0

	for num := range numChan {
		numbers = append(numbers, num)
		count++

		if count == 10 {
			sum := 0
			for _, n := range numbers {
				sum += n
			}
			average := float32(sum) / 10.0
			avrChan <- average
			numbers = numbers[:0]
			count = 0
		}
	}
	close(avrChan)
}
