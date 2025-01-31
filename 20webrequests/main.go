package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Web Request")
	response, err := http.Get(url)

	if err!= nil {
		panic(err);
	}

	fmt.Printf("Status: %s\n", response.Status);
	fmt.Println(response.Body)

	defer response.Body.Close();

	databytes,err := ioutil.ReadAll(response.Body);
	if err!= nil {
        panic(err);
    }

	fmt.Println("Bytes",string(databytes));

}
