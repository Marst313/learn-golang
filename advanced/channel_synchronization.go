package main

import (
	"fmt"
	"time"
)

/* func main() {

	done := make(chan int)
	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- 0
	}()

	fmt.Println(<-done)

	fmt.Println("Finished")

}
*/

/* func main() {
	ch := make(chan int)

	go func() {
		ch <- 9 // Blocking until the value is received
		fmt.Println("Sent value")

	}()

	value := <-ch
	fmt.Println(value)

}
*/
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(2 * time.Second)

// 			done <- id
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done
// 	}

// 	fmt.Println("All goroutine are complete")
// }

// ========= SYNCHRONIZING DATA EXCHANGE
func channel_synchronization() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "Hello" + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	// This loop create receiver each time and stop creating receiveer // ! if channel is closed
	for value := range data {
		fmt.Println("Received value : ", value, ":", time.Now())
	}
}
