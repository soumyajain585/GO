package Controllers

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mu      sync.Mutex // Create a Mutex
)

func Counter() {
	var wg sync.WaitGroup

	// Start multiple goroutines to increment the counter concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock() // Lock the Mutex before accessing the shared variable
				counter++
				mu.Unlock() // Unlock the Mutex when done
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("Final Counter Value: %d\n", counter)
}
