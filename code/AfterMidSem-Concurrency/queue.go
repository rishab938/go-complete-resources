package queue

import (
	"fmt"
)

func main(){

	queue:= make(chan string, 2)

	queue <- "one"
	queue <- "two"

	// fmt.Println(<-queue)
	// fmt.Println(<-queue)

	close(queue)

	for ele:= range queue {
		fmt.Println(ele)
	}

}