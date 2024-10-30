package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	tests := []struct {
		name          string
		numIncrements int
		wantCount     int
	}{
		{
			name:          "5 increments",
			numIncrements: 5,
			wantCount:     5,
		},
		{
			name:          "0 increments",
			numIncrements: 0,
			wantCount:     0,
		},
		{
			name:          "-1 increments",
			numIncrements: -1,
			wantCount:     0,
		},
		{
			name:          "1000 increments",
			numIncrements: 1000,
			wantCount:     1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := sync.WaitGroup{}
			counter := &Counter{}

			for i := 1; i <= tt.numIncrements; i++ {
				wg.Add(1)
				go func(id int) {
					Inc(id, counter)
					wg.Done()
				}(i)
			}
			wg.Wait()
			assert.Equal(t, tt.wantCount, counter.count)
		})
	}
}

func TestStartWorkers(t *testing.T) {
	tests := []struct {
		name        string
		numWorkers  int
		wantWorkers int
		wantErr     error
	}{
		{
			name:        "5 workers",
			numWorkers:  5,
			wantWorkers: 5,
			wantErr:     nil,
		},
		{
			name:        "negative numbers of workers",
			numWorkers:  -1,
			wantWorkers: 0,
			wantErr:     assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := &Counter{}
			wg := sync.WaitGroup{}
			err := startWorkers(tt.numWorkers, counter, &wg)
			if tt.wantErr != nil {
				assert.Error(t, err)
				return
			}
			wg.Wait()
			assert.Equal(t, tt.wantWorkers, counter.count)
		})
	}
}
