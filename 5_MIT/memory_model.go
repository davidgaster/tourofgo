package main

// This program starts a goroutine for every entry in the work list, 
// but the goroutines coordinate using the limit channel to ensure 
// that at most three are running work functions at a time.

var limit = make(chan int, 3)

func main() {
	for _, w := range work { // work is a list of worker routines
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	select{}
}