package main

import (
	"github.com/goodelias/otusgo-basic/hw10_motion_sensor/calc"
)

func main() {
	numChannel := make(chan int)
	go calc.GenerateNumbers(numChannel)
	calc.CalculateAverage(numChannel)
}
