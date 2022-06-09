/*
Channels giup giai quyet cac tai
nguyen dung chung giua cac GoRoutine

*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")
	var wg = sync.WaitGroup{}

	ch := make(chan int)
	// goRoutine cho den khi co' so
	// goRoutine chuyen nhan
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	// goRoutine dua con so vao channels
	// goRoutine chuyen gui

	go func(ch chan<- int) {
		ch <- 42
		ch <- 42
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
