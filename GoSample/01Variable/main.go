/*
- Dinh nghia bien trong ngon ngu lap trinh
- Cu phap khai bao bien: 3 cach
- Globals and block scope
  + global scope thi khai bao
  ngoai than ham va co tu khoa var
  + co the khao bao nhieu var trong cung var {}
- shadow
- Declared and not used
- Export global scope
	+ ky tu dau tien viet hoa.
- Naming convention
- convert type
	+ Doi voi Go thi chuyen so thanh chuoi thi se chuyen sang ma ascii
	=> muon chuyen chinh xac thi phai dung
	var s string = strconv.itoa(10)
*/
package main

import "fmt"

/*
	// Export global scope
	var NumGlobal int =10

	// Globals and block scope
	var {
		n int = 10
		m float64 = 100.0
		k string = "hello string"
	}
*/
func main() {
	var i int
	i = 10
	// var i int = 10.0
	// or
	// i := 10
	fmt.Println(i)
	fmt.Printf("%v, %T", i, i)
}
