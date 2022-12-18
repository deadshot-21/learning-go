package main

import "fmt"

// Capitalzing will creat public var
const LoginToken = "avbjdvajhdvaj"

func main() {
	var username string = "deadshot"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFloat float64 = 255.59624863589952
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n",smallFloat)

	//default value 0
	var someNum int
	fmt.Println(someNum)
	fmt.Printf("Variable is of type: %T \n",someNum)

	//implicit style
	var email = "sohamfaldu@gmail.com"
	fmt.Println(email)
	fmt.Printf("Variable is of type: %T \n",email)

	//no var style
	noOfUsers := 30000.0
	fmt.Println(noOfUsers)
	fmt.Printf("Variable is of type: %T \n",noOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n",LoginToken)

}