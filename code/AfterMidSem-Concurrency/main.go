package main

import (
	"fmt"
	"time"
)


// func greet(name string, ch chan bool){
// 	fmt.Println("Hello", name)
// 	ch <- true // sending value
// }

func greet(name string, ch chan bool){
	start:= time.Now()
	fmt.Println("Hello", name, start.Format("15.04.05.000"))
	time.Sleep(1 * time.Second)
	ch <- true // sender code: blocking
}

func main(){
	fmt.Println("welcome back!!!")

	// ch1:= make(chan bool)
	// ch2:= make(chan bool)
	// ch3:= make(chan bool)

	ch:= make(chan bool) // unbuffered

	go greet("Ravi", ch)
	go greet("ABC", ch)
	<-ch
	<-ch // receiver code: blocking // blocking
	go greet("Yug", ch)
	go greet("Rohit", ch)
	<-ch
	<-ch
	

	// goroutine
	// channel

	// time.Sleep(1 * time.Second)
}