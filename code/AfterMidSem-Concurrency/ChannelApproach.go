package main

import "fmt"

func increment(ch chan int){
	ch <- 1
}


func main(){

	ch:= make(chan int) // channel - unbuffered

	for i:=0; i<5; i++{
		go increment(ch)
	}

	counter:=0

	for i:=0; i<5; i++{
		counter = counter + <-ch
	}

	fmt.Println("Channel Approach:",counter)

}
