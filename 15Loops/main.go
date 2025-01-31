package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")
	days := []string{"Sun", "Mon", "Tue","Wed", "Thu", "Fri", "Sat"}
	
	for i, day := range days {
        fmt.Printf("Day %d: %s\n", i+1, day)
    }

	// for d:=0; d<len(days); d++ {
	// 	fmt.Printf("Day %d: %s\n", d+1, days[d])
	// }

	


}
