package main

import (
	"fmt"
	"time"
)

type Task func()

func main() {
	taskC := make(chan Task)

	for i := 0; i < 4; i++ {
		go func() {
			for task := range taskC {
				task()
			}
		}()
	}

	task := func() {
		fmt.Println("this is test")
	}
	for i := 0; i < 10; i++ {
		taskC <- task
	}

	close(taskC)

	time.Sleep(time.Second)
}
