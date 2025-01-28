package main

import "fmt"

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)

	languages["Go"] = "Google"
	languages["Python"] = "Python Software Foundation"

	fmt.Println(languages)

	fmt.Println(languages["Go"]);

	delete(languages, "Go");

	// loops in map
	for key,value := range languages {
		fmt.Printf("For key %v value %v", key, value);
	}
}
