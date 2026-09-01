package main

import (
	"fmt"
	"sync"
	"time"
)

// CONSTRUCTION EXAMPLE
type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func wait_groups() {
	var wg sync.WaitGroup

	// DEFINE TASKS TO BE PERFORMED BY WORKERS
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	// WAIT FOR ALL WORKERS TO FINISH
	wg.Wait()

	// CONSTRUCTION IS FINISHED
	fmt.Println("Construction finished")

}

/* // ? EXAMPLE WITH CHANNELS
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d starting. \n", id)
	time.Sleep(time.Second) // simulate some work

	for task := range tasks {
		results <- task * 2
	}

	fmt.Printf("WorkerID %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 3
	numJobs := 3
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
*/

/* // ? BASIC EXAMPLE WITHOUT USING CHANNELS
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ! ALWAYS DO THIS

	 wg.Add(1) // ! WRONG PRACTICE

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some time spent on processing the task
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAIT GROUPS

	! LAUNCH WORKERS
	for i := range numWorkers {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
*/
