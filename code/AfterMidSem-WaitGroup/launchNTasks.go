package main

import (
	"fmt"
	"sync"
)

func doTask(id int, wg *sync.WaitGroup){
	//
	defer wg.Done()
	fmt.Printf("Task %d is done\n", id)
	//
	
}

func main(){
	var wg sync.WaitGroup
	for i:=1; i<=3; i++ {
		wg.Add(1)
		doTask(i, &wg)
	}

	wg.Wait()
}