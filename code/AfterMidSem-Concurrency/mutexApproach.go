package main

import (
	"fmt"
	"sync"

)

var counter = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main(){
	var wg sync.WaitGroup

	for i:=0; i<5; i++{
		wg.Add(1)
		go increment(&wg)
	}

	// time.Sleep(1*time.Second)
	wg.Wait()
	fmt.Println(counter)
}