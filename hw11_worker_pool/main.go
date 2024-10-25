package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func Inc(id int, c *Counter) {
	c.mu.Lock()
	c.count++
	fmt.Printf("Worker %d incremented the count to %d\n", id, c.count)
	c.mu.Unlock()
}

func worker(id int, c *Counter) {
	Inc(id, c)
}

func startWorkers(numWorkers int, c *Counter, wg *sync.WaitGroup) error {
	if numWorkers < 0 {
		return fmt.Errorf("number of workers cannot be negative: %d", numWorkers)
	}
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id, c)
			wg.Done()
		}(i)
	}
	return nil
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}

	err := startWorkers(-2, counter, &wg)
	if err != nil {
		fmt.Printf("Failed to start worker: %v", err)
		return
	}

	wg.Wait()
}
