package main

import (
	"sync"
	"time"
	"fmt"
)

// locks are meant to protect ** shared data **
// but they are also meant to protect ** invariance ** i.e.  alice + bob = total
func main() {
	alice := 10000
	bob := 10000
	var mu sync.Mutex

	total := alice + bob

	go func() {
		for i := 0; i < 1000; i++ {
			// we have made increment atomic AND decrement atomic
			// but what we really want is for the entire transaction to be atomic
			mu.Lock()
			alice -= 1
			mu.Unlock()
			mu.Lock()
			bob += 1
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			mu.Unlock()
			mu.Lock()
			alice += 1
			mu.Unlock()
		}
	}()

	// audit that these locks are working correctly, i.e. the total is always the same
	// we see that alice + bob is not equal to the overall sum
	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice + bob != total {
			fmt.Printf("observed violation, alice = %v, bob=%v, sum=%v\n", alice, bob, alice+bob)
		}
		mu.Unlock()
	}
	

}