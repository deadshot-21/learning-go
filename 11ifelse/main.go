package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular size"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 != 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

}
