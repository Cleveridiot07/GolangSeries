package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("Switch Case in Golang")

    rand.Seed(time.Now().UnixNano())
    randomNumber := rand.Intn(10) + 5
    fmt.Printf("Random Number: %d\n", randomNumber)

	switch randomNumber {
		case 5:
            fmt.Println("Number is 5")
        case 6, 7:
            fmt.Println("Number is either 6 or 7")
			fallthrough
        default:
            fmt.Println("Number is neither 5 nor 6 or 7")
			
	}

	// Fallthrough will also print default values because its is wriiten after the switch case
}