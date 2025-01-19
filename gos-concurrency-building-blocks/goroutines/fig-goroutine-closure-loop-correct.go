package main

import (
	"fmt"
	"sync"
)

// This will works fine on Go.v1.22 (range notes part)
// See docs for details : https://tip.golang.org/doc/go1.22
func main() {
	var wg sync.WaitGroup
	// This code will not cause a race condition, since `salutation` is a new variable for each iteration of the range loop.
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
