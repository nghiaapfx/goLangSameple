package main

import (
	"fmt"
)

func main() {
	fmt.Println("Function")

	sum := findSum("hello world", 1, 3, 5, 7, 9)
	fmt.Println(sum)

	res, err := CalDivide(4, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	/// goi method cua 1 struct
	s := Student{name: "Bill Gates"}
	s.AddNewStudent()
	fmt.Println(s)
}

func noReturn(a int, b int) {
	fmt.Println(a, b)
}

func returnValue(a int, b int) (initValue int) {
	initValue = 2

	return
}

func returnValue2(a int, b int) (int, string) {
	initValue := 2
	err := "error message"

	return initValue, err
}

// Khai niem rest trong ES6
func findSum(s string, rest ...int) (sum int) {
	for _, v := range rest {
		sum += v
	}
	fmt.Println(s)
	return
}

func CalDivide(a, b int) (int, error) {
	result := 0
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result = a / b
	return result, nil
}

// method to student
func (obj Student) AddNewStudent() {
	fmt.Println("")
}

type Student struct {
	name string
}
