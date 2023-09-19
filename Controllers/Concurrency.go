package Controllers

import (
	"fmt"
	"sync"
)

func calculateSum(slice []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup when this goroutine is done

	sum := 0
	for _, num := range slice {
		sum += num
	}

	resultChan <- sum // Send the sum to the result channel
}

func Concurrency() {
	// Create two slices of numbers
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to receive results
	resultChan := make(chan int)

	// Launch two goroutines to calculate the sum of each slice concurrently
	wg.Add(2) // Add 2 to the WaitGroup to track two goroutines
	go calculateSum(slice1, resultChan, &wg)
	go calculateSum(slice2, resultChan, &wg)

	// Start a goroutine to close the result channel when all calculations are done
	go func() {
		wg.Wait()         // Wait for both goroutines to finish
		close(resultChan) // Close the result channel to signal that all results are sent
	}()
	fmt.Println("resultChan ", resultChan)
	// Collect and sum the results from the channel
	total := 0
	for sum := range resultChan {
		total += sum
	}

	fmt.Printf("Total Sum: %d\n", total)
}
