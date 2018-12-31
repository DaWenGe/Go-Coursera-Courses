package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	sharedData := []int{0, 0, 0, 0}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		i := rand.Intn(4)
		sharedData[i]++
		wg.Done()
	}()

	go func() {
		j := rand.Intn(4)
		sharedData[j]++
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(sharedData)
}

// Race Condition:
// It happens when two goroutines access the same sharedData at the same time and they are accessing the same item of this sharedData while at least one of them is modifying that item
// Communication is the source of race condition
