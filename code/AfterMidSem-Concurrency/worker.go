package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int){
	fmt.Printf("Execution of %d worker started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Execution of %d worker ended\n", id)

}
func main(){
	var wg sync.WaitGroup
	for i:=1; i<=5; i++{
		wg.Go(func ()  {
			worker(i)
		})
		// go worker(i)
	}

	wg.Wait()
}