package main

import "fmt"

func multiplexing_select() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received: ", msg)
		}
	}

}

/* func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received from ch 1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch 2:", msg)
		// default:
		// fmt.Println("No channels ready...")
	}
	fmt.Println("End of program")

} */
