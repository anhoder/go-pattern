package main

import (
	"fmt"
	"time"
)

func main() {
	finishC := make(chan struct{})

	go func() {
		time.Sleep(time.Millisecond * 200)

		close(finishC)
	}()

	res, ok := <-finishC
	fmt.Println(res, ok)
}
