package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	rand.Seed(time.Now().UnixMicro())

	go func() {
		r := rand.Intn(500)
		fmt.Println("rand:", r)
		time.Sleep(time.Duration(r) * time.Millisecond)
		ch <- "this is data"
	}()

	tc := time.After(time.Millisecond * 200)
	select {
	case data := <-ch:
		fmt.Println(data)
	case <-tc:
		fmt.Println("time out")
	}
}
