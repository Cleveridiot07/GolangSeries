package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	user := User{Name: "John Doe", Email: "johndoe@example.com", Status: true, Age: 30}

	user2 := User{"John", "John", false, 50}

	fmt.Printf("User: %+v\n", user)
	fmt.Printf("User 2 %v\n", user2)

	fmt.Printf("User Name is %v\n", user.Name);


	user2.GetStatus();
	user2.NewEmail();

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


// Example Getter
func(U User) GetStatus(){
	fmt.Println("is User Active: ", U.Status);
}

// Example Setter But it doesn't change the value pf U passed as it only change the copied U
func (U User) NewEmail(){
	U.Email = "Test@go.dev"
	fmt.Println("Updated Email: ", U.Email);

}

// func (U User) NewEmail(){
// 	U.Email = "Test@go.dev"
// 	fmt.Println("Updated Email: ", U.Email);

// }
