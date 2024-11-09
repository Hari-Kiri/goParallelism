package main

import (
	"fmt"
	"runtime"
	"sync"
)

const processor = 2

func main() {
	runtime.GOMAXPROCS(processor)

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go func() {
		defer waitGroup.Done()
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()
	go func() {
		defer waitGroup.Done()
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()
	waitGroup.Wait()
}
