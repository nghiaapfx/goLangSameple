/*


 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pipeline -fan-in fan-out")
	// generate Pipeline
	randNumber := []int{}
	for i := 0; i < 1000000000; i++ {
		randNumber = append(randNumber, i)
	}
	inputChan := generatePipeline(randNumber)

	//fan out
	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	//fan In
	c := fanIn(c1, c2, c3, c4)
	sum := 0
	for i := 0; i < len(randNumber); i++ {
		sum += <-c
	}
	fmt.Printf(" Total %d", sum)
}
func generatePipeline(number []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range number {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChan ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChan {
			for n := range c {
				in <- n
			}
		}
	}()
	return in
}
