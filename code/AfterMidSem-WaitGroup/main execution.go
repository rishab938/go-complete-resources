package main

import "fmt"

func greet(str string){
	fmt.Println("Hello", str)
}

func main(){

	fmt.Println("Hello")
	go greet("ABC")

}