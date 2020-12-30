package main

import "sync"

func main() {
	println("good goroutine")
	good()
	println("bad goroutine")
	bad()
}

/*
	Go pattern: when using a goroutine in a loop, make sure to pass in the
	i as an argument to the go-routine, otherwise it will take the i that's being
	mutated in the for loop.
*/
func good() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			sendRPC(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func bad() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			sendRPC(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func sendRPC(i int) {
	println(i)
}