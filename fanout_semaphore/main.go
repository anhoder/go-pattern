package main

import (
	"fmt"
	"time"
)

func main() {
	num := 50
	ch := make(chan string, num)

	tokenC := make(chan struct{}, 5)

	for i := 0; i < num; i++ {
		go func(no int) {
			tokenC <- struct{}{}
			{
				ch <- fmt.Sprintf("Hello, this is %d", no)
				time.Sleep(500 * time.Millisecond)
			}
			<-tokenC
		}(i)
	}

	for ; num > 0; num-- {
		fmt.Println(<-ch)
	}
}
