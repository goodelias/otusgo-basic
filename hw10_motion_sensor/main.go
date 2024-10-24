package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/goodelias/otusgo-basic/hw10_motion_sensor/calc"
)

func main() {
	in := make(chan int)
	out := make(chan float32)

	// goroutine for generating numbers
	go func() {
		timer := time.NewTimer(time.Minute)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				num := rand.Intn(10) //#nosec
				fmt.Printf("Generated new number: %d\n", num)
				in <- num
			case <-timer.C:
				close(in)
				fmt.Println("Time is up. Generation stopped")
				return
			}
		}
	}()

	go calc.CalculateAverage(in, out)

	func() {
		for average := range out {
			fmt.Printf("The arithmetic mean for the last numbers is %.2f\n", average)
		}
	}()
}
