package calc

func CalculateAverage(in <-chan int, out chan<- float32) {
	defer close(out)
	numbers := make([]int, 0, 10)
	count := 0

	for num := range in {
		numbers = append(numbers, num)
		count++

		if count == 10 {
			sum := 0
			for _, n := range numbers {
				sum += n
			}
			average := float32(sum) / 10.0
			out <- average
			numbers = numbers[:0]
			count = 0
		}
	}

	if count > 0 {
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		average := float32(sum) / float32(count)
		out <- average
	}
}
