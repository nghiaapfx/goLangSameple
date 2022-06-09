/*
Dinh nghia hang so
Khai bao hang so
Dac diem cua hang so
Kieu du lieu enum
Tu dong khoi tao gia tri cho enum bang iota
*/

package main

import (
	"fmt"
)

const i = 42

//enum
const (
	red    = 0
	yellow = 1
	blue   = 2
	black  = 3
)

// default iota chay tu 0 den 4

const (
	_ = iota + 5
	pink
	white
)

func main() {
	fmt.Println(" Constants - Enum ")

	// hang so se bi shadow (override)
	// trong pham vi ham
	// output: i= 12
	const i = 12
	fmt.Printf("%v, %T", i, i)

	// enum
	fmt.Printf(" enum: %v, %T", red, red)

	// iota
	fmt.Printf(" iota: %v, %T", pink, pink)
}
