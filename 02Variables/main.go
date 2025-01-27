package main

import "fmt"

// jwtTokeb := 30000 var less operator are not supported outside functions or gloabal

// variable with firsr letter as capital is a public variable
const LoginToken string= "login"

func main() {
	// fmt.Println("Variables")
	var username string = "clever"
	fmt.Printf("variable type %T \n", username)

	var is bool = true
	fmt.Printf("variable type %T \n", is)

	var smallVal uint = 256
	fmt.Printf("variable type %T \n", smallVal)

	var floatVal float32 = 1024.455546468665456465464546
	// fmt.Printf("variable  \n", floatVal);
	fmt.Printf("variable type %T \n", floatVal)

	// default value for variable type integer
	var anotherVal int
	fmt.Println(anotherVal)

	// implicit value for variable type
	var website = "http://cleveridiot.netlify.app"
	fmt.Println(website);

	// no var type
	numberOfUser := 6000
	fmt.Println(numberOfUser);


}
