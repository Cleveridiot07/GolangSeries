package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Welcome to time")
    presentTime := time.Now()
    fmt.Println(presentTime)
    fmt.Println(presentTime.Format("2006-01-02 Tuesday 15:04:05 "))



	createdDate := time.Date(2020,time.January,20,23,24,0,0, time.UTC);
	fmt.Println(createdDate);

	fmt.Println(createdDate.Format("2006-01-02 Monday"))
}