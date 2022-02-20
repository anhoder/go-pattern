package main

import "fmt"

func main() {
	num := 50
	ch := make(chan string, num)

	for i := 0; i < num; i++ {
		go func(no int) {
			ch <- fmt.Sprintf("Hello, this is %d", no)
		}(i)
	}

	for ; num > 0; num-- {
		fmt.Println(<-ch)
	}
}
