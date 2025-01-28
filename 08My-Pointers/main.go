package main;

import "fmt"

func main(){
	fmt.Println("Pointers")

	var one int = 2;
	var ptr *int = &one;
	fmt.Println("Address of one: ", ptr);
	fmt.Println("value: ", *ptr);

	*ptr = *ptr*3;
	fmt.Println("value after multiplication: ", *ptr);
}