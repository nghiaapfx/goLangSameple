/*
map la tap hop
cac gia tri anh xa key: value
*/
package main

import (
	"fmt"
	"reflect"
)

type Department struct {
	departmentId   int
	departmentName string
	// tag `Maximum cannot over 50`
	//departmentNameTag string `Maximum cannot over 50`
}
type Student struct {
	numId   int
	name    string
	isMale  bool
	subject []string
	Department
}

func main() {
	fmt.Println(" Map - Struct ")
	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int{
		"Tomato": 10,
		"Carrot": 15,
		"Apple":  20,
	}
	//add student
	studentNameAgeMap["Microsoft"] = 25
	//update student
	studentNameAgeMap["Carrot"] = 16
	//delete student
	delete(studentNameAgeMap, "Tomato")
	// check exist
	_, isExist := studentNameAgeMap["Tomato"]
	copyMap := studentNameAgeMap

	//Struct student
	studentCamel := Student{
		numId:  16,
		name:   "Camel Nguyen",
		isMale: false,
		subject: []string{
			"Math",
			"English",
			"Chemist"},
	}

	studentNewComer := Student{}
	studentNewComer.name = "No Name"
	studentNewComer.departmentName = "R&D class"

	reflectParams := reflect.TypeOf(studentNewComer)

	fmt.Println(" Map ", studentNameAgeMap)
	fmt.Println(" isExist", isExist)
	fmt.Println("copyMap", copyMap)

	fmt.Println(" Struct studentCamel: ", studentCamel.name)
	fmt.Println(" Struct studentNewComer: ", studentNewComer.name)
	fmt.Println("reflect", reflectParams)

}
