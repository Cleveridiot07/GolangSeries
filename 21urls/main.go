package main

import (
	"fmt"
	"net/url"
)

// const url string = "https://catfact.ninja/fact"
const myurl string = "http://localhost:8000/learn?coursename=reactjs&paymentid=ghsyte45"

func main() {
	fmt.Println("Handling Urls in Golang")
	fmt.Println("URL: ", myurl)

	// Parsing URL
	result,err := url.Parse(myurl);
	if err!=nil {
        fmt.Println("Error parsing URL: ", err)
    }
	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Raw Query: ", result.RawQuery)

	qparams := result.Query()
	fmt.Println("Query Parameters:" ,qparams); 

	for val,_ := range qparams {
		fmt.Printf(val);
	}



	partsOfUrl := &url.URL{
		Scheme: "http",
		Host: "test.dev",
		Path:"test",
		RawPath: "user=clever",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println("New URL: ", anotherURL);
	// fmt.Println("coursename: ", qparams[coursename]);
}
