package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer class")

	myNumber:=23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer: ",ptr)
	fmt.Println("Value: ",*ptr)

}