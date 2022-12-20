package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List is: ",fruitList)

	var vegList = [5]string{"Potato","beans","mushrooms"}

	fmt.Println("Veg List is: ",vegList)
	fmt.Println("Veg List length is : ",len(vegList))


}
