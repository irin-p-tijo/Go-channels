package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan func(), 10)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for f := range ch {

				f()
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		ch <- func(i int) func() {
			return func() {
				fmt.Println("Work -", i)
			}

		}(i)
	}
	close(ch)
	wg.Wait()
}
