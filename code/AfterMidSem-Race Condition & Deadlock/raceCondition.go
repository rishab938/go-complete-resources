package main

import (
	"fmt"
	"sync"
)

var balance = 1000
var mu sync.Mutex

func deduct(amount int, wg *sync.WaitGroup){
	defer wg.Done()

	
	mu.Lock()
	if (balance >= amount){
		balance = balance - amount// read and write

	}
	mu.Unlock()

}

func main(){

	amount:= 10
	var wg sync.WaitGroup
	for i:=1; i<=100; i++{
		wg.Add(1)
		go deduct(amount, &wg)
	}

	wg.Wait()

	fmt.Println("Expected Balance ", 0)
	fmt.Println("Actual Balance is ", balance)

	if balance != 0 {
		fmt.Println("Race Condition exists")
	}

}