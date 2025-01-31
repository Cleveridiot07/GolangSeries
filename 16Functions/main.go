package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang");
	greeter();

	result := adder(5,6);
	fmt.Println("Addition is ",result);

	// var arr = []int{1,2,3,4,5};

	result2 := proAdder(5,8,7,6,2,5,4,1);
	fmt.Println("Procedural addition is ",result2);
	
	a, b := retTwoValues()
	fmt.Println("Returned values are: ",a, b);

	// Functions definitions are not allowed in another fuctions
	// func greeterTwo(){
	// 	fmt.Println("Hello, Greeter Two!");
	// }
	// greeterTwo();
}

func adder(a int, b int) int {
    return a + b;
}

func proAdder(values ...int)int{
	total := 0
	for _, value := range values {
        total += value
    }
	return total;
}

func retTwoValues() (int, string) {
	return 10, "World"
}


func greeter(){
	fmt.Println("Hello, Golang!");
}