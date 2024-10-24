package main

import (
	"fmt"
	"sync"

	"github.com/goodelias/otusgo-basic/hw10_motion_sensor/calc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	numChannel := make(chan int)
	avrChan := make(chan float32)
	go calc.GenerateNumbers(numChannel)
	go calc.CalculateAverage(numChannel, avrChan)

	go func() {
		for average := range avrChan {
			fmt.Printf("The arithmetic mean for the last 10 numbers is %.2f\n", average)
		}
		wg.Done()
	}()

	wg.Wait()
}
