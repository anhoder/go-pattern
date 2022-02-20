package main

import (
	"fmt"
	"time"
)

func main() {
	resultC := make(chan string, 5)

	go func() {
		for res := range resultC {
			fmt.Println(res)
		}
	}()

	for i := 0; i < 100; i++ {
		go func(num int) {
			select {
			case resultC <- fmt.Sprintf("this is %d", num):
			default: // drop data when chan is full
			}
		}(i)
	}

	time.Sleep(time.Second)
	close(resultC)
}
