package main

import (
	"fmt"
	"sync"
)

type Task func()

func main() {
	taskC := make(chan Task)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t := <-taskC
		t()

		wg.Done()
	}()

	taskC <- func() {
		fmt.Println("Hello, this is task")
	}

	wg.Wait()
}
