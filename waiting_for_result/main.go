package main

import "fmt"

func main() {

	resultC := make(chan string)

	go func() {
		// do something...
		resultC <- "Hello, this is result"
	}()

	res := <-resultC
	fmt.Println(res)
}
