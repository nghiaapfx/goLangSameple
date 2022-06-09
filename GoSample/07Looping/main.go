package main

import (
	"fmt"
)

func main() {
	fmt.Println("Looping")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	//////////////////////
	fmt.Print("\n")
	j := 0
	for ; j < 10; j++ {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}
	///////////////////
	fmt.Print("\n")
	k := 0
	for k < 10 {
		if k%2 == 0 {
			fmt.Println(k)
		}
		k++
	}
	////////////////////////
	fmt.Print("\n")
	h := 0
	for {
		if h%2 == 0 {
			fmt.Println(h)
		}
		h++
		if h > 10 {
			break
		}
	}

	/////////////////////////////
	fmt.Print("\n")
	arr := [3]int{1, 5, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//////////////////////////
	fmt.Print("\n")
	for index, value := range arr {
		fmt.Println(index, value)
	}
	///////////////////////////
	// Map se random  - khong theo thu tu
	fmt.Print("\n")
	mapp := map[string]int{
		"Carrot": 19,
		"Tomato": 20,
		"Mint":   21,
		"Golden": 15,
	}
	for index, value := range mapp {
		fmt.Println(index, value)
	}
}
