package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeting()
	fmt.Println(adder(2,3))
	proRes , _ := proAdder(2,3,6,7)
	fmt.Println(proRes)

}

func greeting() {
	fmt.Println("Good Evening!")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int,string){
	total:=0
	for _, val := range values {
		total+=val
	}
	return total, "Done"
}