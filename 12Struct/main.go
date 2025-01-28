package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	user := User{Name: "John Doe", Email: "johndoe@example.com", Status: true, Age: 30}

	user2 := User{"John", "John", false, 50}

	fmt.Printf("User: %+v\n", user)
	fmt.Printf("User 2 %v\n", user2)

	fmt.Printf("User Name is %v\n", user.Name);

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
