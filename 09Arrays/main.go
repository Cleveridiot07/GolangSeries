package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Arrays")
	var list  [4]string;
	list[0] = "Apple"
	list[1] = "Banana"
	list[2] = "Cherry"
	list[3] = "Date"
	fmt.Println(list) 

	fmt.Println("len: ", len(list));

	var vegList [3]string;
	vegList[0] = "Carrot"
	vegList[1] = "Cucumber"
	vegList[2] = "Onion"
	fmt.Println(vegList)
}
