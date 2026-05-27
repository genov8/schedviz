package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	deadline := time.Now().Add(2 * time.Second)

	x := 0
	for time.Now().Before(deadline) {
		for i := 0; i < 1_000_000; i++ {
			x += i
		}
	}
	_ = x
}

func main() {
	const workers = 200

	fmt.Println("starting goroutines...")

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println("all goroutines finished")

	time.Sleep(3 * time.Second)
}
