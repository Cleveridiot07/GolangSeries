package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := "Example Content for file"

	file, err := os.Create("./myFile.txt")
	checkNilErr(err);

	length, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully wrote %d bytes to file\n", length)

	readFile("./myFile.txt")

	defer file.Close()

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
