package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointer")
	var a int = 12
	// b la ban sao cua a
	var b int = a
	fmt.Println(a, b)
	a = 24
	fmt.Println(a, b)

	// khai bao con tro => var c *int
	// pointer c tro ve dia chi cua a
	var c *int = &a

	fmt.Println(a, c)
	// output a,c : 24, 0xc0000c0c0

	fmt.Println(&a, c)
	// output a,c : 0xc0000c0c0, 0xc0000c0c0

	// muon xuat ra gia tri cua con tro C
	// thi dung *c

	fmt.Println(a, *c)
	////////////////////////////
	//thay doi gia tri array thi
	// khong lam thay doi arrB
	print("\n")
	arrA := [3]int{1, 2, 3}
	arrB := arrA
	fmt.Println(arrA, arrB)
	arrA[1] = 10
	fmt.Println(arrA, arrB)
	////////////////
	// ban chat slice la con tro,
	// nen chi luu dia chi con vi vay
	// khi thay doi arrC cung thay doi arrD
	print("\n")
	arrC := []int{1, 2, 3}
	arrD := arrA
	fmt.Println(arrC, arrD)
	arrC[1] = 10
	fmt.Println(arrC, arrD)
	////////////////////////////////
	// Khac voi C /C++ khong nen thay doi dia chi con tro
	// thong qua cac phep toan
	// hoac dung package unsafe
	// de thao tac (+ /-) dia chi con tro

	////////////////////////////////
	// struc pointer
	print("\n")
	// con tro struct co gia tri khoi tao la <nil>
	var pointerStruct *mystruct
	fmt.Println(pointerStruct)
	// khi cap phat dung nho se cap
	// con tro 1 gia tri mac dinh
	pointerStruct = new(mystruct)
	fmt.Println(pointerStruct)

	//cap nhat gia tri cho con tro.
	// nen su dung cach nay
	//  => (*pointerStruct).num= 20
	pointerStruct.num = 20
	fmt.Println(pointerStruct.num)

	//output: <nil>
	//        &{0}
	//        20

	/////////////
	// Khi tao tac con tro tren Map
	// cung giong nhu tren slice
}

type mystruct struct {
	num int
}
