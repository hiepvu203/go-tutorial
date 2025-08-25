package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Alo")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go say("world", &wg)
	go say("hello", &wg) // chạy song song
	go hello(&wg)

	fmt.Println("main function")
	wg.Wait()
}
