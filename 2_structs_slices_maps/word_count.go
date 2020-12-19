package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _,word := range strings.Fields(s) {
		wc[word] += 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
