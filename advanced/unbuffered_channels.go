package main

import (
	"fmt"
)

func unbufferred() {
	ch := make(chan int) // imidiately declare channel without maximum memory

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
}
