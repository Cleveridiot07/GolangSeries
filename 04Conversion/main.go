package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	fmt.Println("Please provide a rating between 0 and 10")

	reader := bufio.NewReader(os.Stdin)
	ratingInput, _ := reader.ReadString('\n')

	num,err := strconv.ParseFloat(strings.TrimSpace(ratingInput),64);

	fmt.Println("Thanks for rating", num);

	if(err != nil){
		fmt.Println("Invalid input. Please provide a number between 0 and 10")
	}else{
		fmt.Println("added 1", num+1);
	}

}
