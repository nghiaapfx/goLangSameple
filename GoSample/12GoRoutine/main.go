/*
Goroutines giup tao ra 1 tien trinh doc lap
chay dong thoi voi cac tien trinh khac
- Trong Golang, Khi goroutines cua ham main ket thuc
 thi se stop luon chuong trinh,
  mac ke cac Goroutines khac co thuc hien xong hay khong.

 - WaitGroup: giup quan ly goRoutine va cho den khi
 thuc hien xong.
  - Mutex : la cac lock - giup dong bo hoa cac loi goi ham
	=> khuyet diem : ko tan dung xu ly dong thoi
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("GoRoutine")

	/*
		go count("sheep")
		count("fish")
		time.Sleep(time.Second * 2)
	*/
	var wg = sync.WaitGroup{}

	//so luong GoRoutine phai cho``
	wg.Add(2)

	go func() {
		count("sheep")
		wg.Done()
	}()
	go func() {
		count("fish")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Done")
	////////////////////////////////

	for i := 0; i < 10; i++ {
		wgg.Add(2)
		mux.RLock()
		go sayHello()
		mux.Lock()
		go increment()
	}
	wgg.Wait()
}
func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

var wgg = sync.WaitGroup{}
var counter = 0
var mux = sync.RWMutex{}

func sayHello() {
	fmt.Printf("Hello %v", counter)
	mux.RUnlock()
	wgg.Done()
}

func increment() {
	counter++
	mux.Unlock()
	wgg.Done()
}
