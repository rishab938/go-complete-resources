package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for job:= range jobs {
		fmt.Printf("Worker %d executing job %d\n", id, job)
		time.Sleep(1 * time.Second)
	}

}

func main(){

	const numWorkers = 3
	const numberJobs = 10

	jobs:= make(chan int, 10)

	var wg sync.WaitGroup

	// loop
	for i:=1; i<=numWorkers; i++{
		wg.Add(1) // 1 or i?
		go worker(i, jobs, &wg)
	}

	// jobs loop

	for i:=1; i<=numberJobs; i++{
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	fmt.Println("Jobs executed")
}