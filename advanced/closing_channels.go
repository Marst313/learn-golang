package main

import "fmt"

func producer_closing(ch chan<- int) {
	for i := range 10 {
		ch <- i
		fmt.Println(i, "< i ")
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func closing_channels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer_closing(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val, "ini genap")

	}

}

// ! AVOID MULTIPLE CLOSING CHANNEL
/*
func main() {
	ch := make(chan int)

	go func() {
		close(ch)
		close(ch) // !!! Dont do this
	}()

	time.Sleep(2 * time.Second)
}
*/

// ! RANGE OVER CLOSED CHANNEL
/* func main() {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received: ", val)
	}
}
*/
// ! RECEIVING FROM A CLOSED CHANNEL
/* func main() {
	ch := make(chan int)
	close(ch)

	val, ok := <-ch

	if !ok {
		fmt.Println("Channel is closed")
		return
	}
	fmt.Println(val)
}
*/

// ! SIMPLE CLOSING CHANNEL EXAMPLE
/*
func main() {
	ch := make(chan int)

	go func() {
		ch <- 13
		close(ch)
	}()

	fmt.Println(<-ch, "Closing simpel channel")
}
*/
