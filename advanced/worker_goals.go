package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)

		// simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

func worker_goals() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor/worker

	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!", <-ticketResults)

	}

}

// ! BASIC WORKER POOL PATTERN
/*
func worker(id int, tasks <-chan int, results chan<- int) {
	for {
		select {
		case task := <-tasks:
			fmt.Printf("Worker %d: processing task %d\n", id+1, task)
			time.Sleep(time.Second)
			results <- task
		default:
			fmt.Printf("Worker %d: nganggur \n", id+1)
		}

	}
}

func main() {
	numWorkers := 3
	numJobs := 20
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	!  CREATE WORKERS
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	! Send Values to the tasks Channel
	for i := range numJobs {
		tasks <- i
		fmt.Printf("Send Value to tasks %d\n", i+1)

	}

	! Collect the results
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}

}
*/
