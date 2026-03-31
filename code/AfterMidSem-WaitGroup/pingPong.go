package main

import "fmt"

// This ping function only accepts a channel for sending values. 
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
    pings <- msg
}
// The pong function accepts one channel for receives (pings) and a second for sends (pongs).

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}


// send only channel
func send(ch chan<- int){
	ch<-3
	// val:= <-ch
}

// receive only channel
func receive(ch <-chan int){
	val:=<-ch
	fmt.Println(val)
}

func main(){

	ch:= make(chan int, 1)
	send(ch)
	fmt.Println(<-ch)
	
	ch2:= make(chan int, 1)
	ch2<-10
	receive(ch2)


	pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}