package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter your name: ")
	

	// comma ok syntax || error ok
	// eother you will recieve an input or an error
	name ,_:= reader.ReadString('\n');
	fmt.Printf("Hello, %s", name);

	// if you want to handle the error you can use the following syntax
	name2 ,err:= reader.ReadString('\n');
	fmt.Printf("Hello, %s", name2);
	if err != nil {
        fmt.Println("Error reading input:", err)
    }

}
