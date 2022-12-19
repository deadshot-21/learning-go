package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Ratings: ")

	// comma,err syntax

	ratings, _ := reader.ReadString('\n')
	fmt.Println(ratings)
}
