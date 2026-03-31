package main

import (
	"fmt"
	"sync"
	// "time"
)

func printN(n int, wg *sync.WaitGroup){
	// 2
	for i:=1; i<=n; i++{
		fmt.Println(i)
	}
	// 3 
	wg.Done()
}

func main(){
	const num = 5
	var wg sync.WaitGroup
	wg.Add(1)
	go printN(num, &wg)
	// 1
	// time.Sleep(1 * time.Second)
	wg.Wait()
}