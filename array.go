package main

import "fmt"

func main() {

	// var myMap map[string]int
	var myArr [3]int

	// myMap = make(map[int]int)
	// myMap := make(map[string]int)
	myMap := map[string]int{}
	myMap["Andika"] = 10
	myMap["0"] = 10
	myMap["1"] = 20
	myMap["2"] = 30

	myArr[0] = 10
	myArr[1] = 20
	myArr[2] = 30

	fmt.Printf("%+v \n", myMap)
	fmt.Printf("%+v \n\n", myArr)

}
