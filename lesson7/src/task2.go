package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	naturalsLimit := 100
	naturals := make(chan int)
	cancel := make(chan bool)

	go func() {
		for x := 0; x <= naturalsLimit; x++ {
			naturals <- x
		}
		wg.Done()
		cancel <- true
		return
	}()

	go func() {
		for {
			select {
			case x := <-naturals:
				fmt.Println(x * x)
			case _ = <-cancel:
				//defer close(naturals)
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
	close(cancel)
}
