package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context){
	fmt.Println("Worker starting")
	// time.Sleep(2*time.Second)
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Worker cancelled")
	}
	fmt.Println("Worker ending")
}

func main(){

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(3*time.Second)

}