package main

import "fmt"

// bidirectional

// bcoz we awant to control the direction
// send only channel
// receive only channel

func send(ch chan<- int){
	ch<-1
}

func receive(ch <-chan int){
	val:=<-ch
	fmt.Println(val)
}

// send only
func ping(ch chan<- int, msg int){
	ch<-msg
}

func pong(ch1 <-chan int, ch2 chan<- int){
	valRec:= <- ch1 // received
	ch2 <- valRec // send
}


func main(){

	ch:= make(chan int, 1)
	send(ch)
	// val:= <-ch

	pings:= make(chan int, 1) 
	pongs:= make(chan int, 2)

	ping(pings, 100)
	pong(pings, pongs)
	fmt.Println(<-pongs)
}