package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	soham := User{"Soham","s@g.com",true,21}
	fmt.Println(soham)
	fmt.Printf("Soham details are %+v\n",soham)
	fmt.Printf("Age is %v",soham.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
