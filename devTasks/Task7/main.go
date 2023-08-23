package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Work Time: %v\n", time.Since(start))
}

func or[T any](channels ...<-chan T) <-chan T {
	res := make(chan T)
	for _, singleChan := range channels {
		go func(ch <-chan T) {
			select {
			case <-ch:
				close(res)
				return
			}
		}(singleChan)
	}
	return res
}
