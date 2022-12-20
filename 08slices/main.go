package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange"}
	fmt.Println("Fruit List: ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit List: ", fruitList)

	fruitList = append(fruitList[0:3])
	fmt.Println("Fruit List: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 218
	highScores[3] = 892

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	// remove value at index 2
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}
