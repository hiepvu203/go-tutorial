package main

import (
	"fmt"
	"time"
)

func sumCal(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sumCal(s[:len(s)/2], c)
	//go sumCal(s[len(s)/2:], c)
	//x, y := <-c, <-c // receive from c
	//
	//fmt.Println(x, y, x+y)

	//ch := make(chan int)
	//
	//go func() {
	//	ch <- 10 // chờ đến khi có goroutine khác nhận
	//}()
	//
	//fmt.Println(<-ch) // 10

	//ch := make(chan int, 2)
	//
	//ch <- 1
	//ch <- 2
	//// ch <- 3 // block vì buffer đầy
	//
	//fmt.Println(<-ch) // 1
	//fmt.Println(<-ch) // 2
	//
	//close(ch)
	//
	//_, ok := <-ch
	//if !ok {
	//	fmt.Println("channel closed")
	//}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "from ch1" }()
	go func() { ch2 <- "from ch2" }()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received", msg2)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
}
