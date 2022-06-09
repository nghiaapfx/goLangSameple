/*
Dung de xu ly cac loi khong mong muon

Defer: Cac lenh phia sau lenh defer se duoc thuc
hien truoc khi ham ben trong no tra ve ket qua
=> defer se thuc hien sau khi ket thuc ham
va chuan bi tra ve ket qua
Defer se dua dong lenh vao ngan xep va xu ly sau

panic: la cac loi xay ra khong mong muon
khi chuong trinh chay
 panic : giup thay doi thoong tin loi
 xuat ra man hinh

recover: kiem tra co panic xay ra hay khong,
neu co thi se dua vao loi bien loi error
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer - panic - Recovery")

	// fmt.Println("Line number 1")
	// defer fmt.Println("Line number 2")
	// fmt.Println("Line number 3")
	// fmt.Println("Line number 4")
	// fmt.Println("Line number 5")
	// output: 1, 3, 4, 5, 2
	////////////////////////////////
	fmt.Print("\n")
	// dua cac lenh co
	//tu khoa defer vao ngan xep

	// defer fmt.Println("Line number 1")
	// defer fmt.Println("Line number 2")
	// defer fmt.Println("Line number 3")
	// defer fmt.Println("Line number 4")
	// defer fmt.Println("Line number 5")

	// output: 5, 4, 3, 2, 1
	///////////////////////////////

	// a := 12
	// defer fmt.Println(a)
	// a = 24
	// output: a = 12
	// vi lenh defer da tao 1 ban sao truoc do

	//////////////////////////////////////////
	// defer fmt.Println("defer")
	// c := 10
	// d := 0
	// panic("Chia cho 1 so 0")
	// fmt.Println(c / d)

	// chuong trinh se dung chay khi gap panic
	// tuong tu lenh throw..
	// defer se thuoc hien truoc khi
	// lenh panic thuc hien

	///////////////////////////
	// cach xu ly truoc khi panic xay ra
	// la dung ham an danh
	// ham an danh chi goi 1 lan duy nhat

	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}
func panicker() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)

		}
	}()
	panic(" chia cho mot so 0")

}
