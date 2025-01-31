package main

import "fmt"

func main() {
	// Defer send the function call at last ie  just before the end of uly braces of parent function if multiple functions are deferred they will execute in last in first out(LIFO) manner
	defer fmt.Println("Hello, World! One")
	defer fmt.Println("Hello, World! Two")
	defer fmt.Println("Hello, World! Three")
	fmt.Println("Test");

	myDefer();
}


func myDefer(){
	for i:=0 ;i <5;i++{
		defer fmt.Println(i);
	}
}