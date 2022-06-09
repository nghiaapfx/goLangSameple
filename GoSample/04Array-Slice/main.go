package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array - Slice")

	//arr := [3]int{1, 3, 7}
	arr := [...]int{1, 3, 7, 9}

	//var arr []int

	// copy array
	brr := arr
	brr[1] = 20

	// khi muon crr tro den dia chi cua arr
	// de tranh lang phi memory thi..
	crr := &arr

	// thay doi crr cung lam thay doi arr
	crr[1] = 30

	// slice
	sliceArr := []int{1, 3, 7, 9, 11, 15}

	sliceBrr := sliceArr[:]
	sliceCrr := sliceArr[3:]
	sliceDrr := sliceArr[:5]
	sliceErr := sliceArr[3:5]

	//tao kich thuoc dong
	sliceMakeArr := make([]int, 0)
	sliceMakeArr = append(sliceMakeArr, 1)
	sliceMakeArr = append(sliceMakeArr, 12, 5, 7)
	//khi them 1 slide thi cuoi phai them dau "..."
	sliceMakeArr = append(sliceMakeArr, []int{20, 30, 40}...)
	fmt.Printf(" Array: %v, %T", arr, arr)
	fmt.Println(" Lenght: ", len(arr))

	fmt.Printf(" brr Array: %v, %T", brr, brr)

	fmt.Printf(" crr Array: %v, %T", crr, crr)

	fmt.Printf(" sliceArr lenght %v, %v ", len(sliceArr), cap(sliceArr))
	fmt.Printf(" sliceArr lenght %v, %v ", len(sliceArr), cap(sliceArr))
	fmt.Printf(" sliceBrr lenght %v, %v ", len(sliceBrr), cap(sliceBrr))
	fmt.Printf(" sliceCrr lenght %v, %v ", len(sliceCrr), cap(sliceCrr))
	fmt.Printf(" sliceDrr lenght %v, %v ", len(sliceDrr), cap(sliceDrr))
	fmt.Printf(" sliceErr lenght %v, %v ", len(sliceErr), cap(sliceErr))
	fmt.Printf(" sliceMakeArr lenght %v, %v ", len(sliceMakeArr), cap(sliceMakeArr))
}
