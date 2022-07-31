package main

import "fmt"

func main() {

	ch := make(chan int)
	//ch <- 2
	//fmt.Println("종료")
	select {
	case ch <- 1:
		fmt.Println("case")
	default:
		fmt.Println("default")
	}
	fmt.Println("종료")
}
