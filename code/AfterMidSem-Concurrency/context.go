package main

import (
	"context"
	"time"
)



func main(){
	ctx:= context.Background()
	ctx2, cancel := context.WithCancel(context.Background())
	ctx3, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
}
