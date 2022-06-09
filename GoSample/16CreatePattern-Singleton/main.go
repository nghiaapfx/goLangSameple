/*

 */
package main

import (
	"fmt"
	"singleton"
)

func main() {
	fmt.Println("Create Pattern - Singleton")
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()

	fmt.Println("%p", s1)
	fmt.Println("%p", s2)
}
