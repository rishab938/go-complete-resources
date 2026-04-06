// PROBLEM: Deadlock via Opposite Lock Ordering
//
// Two goroutines each need to acquire two mutexes (lockA and lockB).
// Goroutine 1 locks A first, then tries to lock B.
// Goroutine 2 locks B first, then tries to lock A.
// This creates a circular wait: each holds one lock the other needs,
// so neither can proceed — resulting in a deadlock.
// Fix: Always acquire locks in the same order (e.g., A → B) in all goroutines.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var lockA, lockB sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// Goroutine 1: locks A then B
	go func() {
		defer wg.Done()
		lockA.Lock()
		fmt.Println("Goroutine 1: locked A, waiting for B...")
		lockB.Lock()
		fmt.Println("Goroutine 1: got both locks")
		lockB.Unlock()
		lockA.Unlock()
	}()

	// Goroutine 2: locks B then A (opposite order = DEADLOCK!)
	go func() {
		defer wg.Done()
		lockB.Lock()
		fmt.Println("Goroutine 2: locked B, waiting for A...")
		lockA.Lock()
		fmt.Println("Goroutine 2: got both locks")
		lockA.Unlock()
		lockB.Unlock()

	}()

	wg.Wait()
	fmt.Println("Done")
}
