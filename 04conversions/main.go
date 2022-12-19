package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Ratings: ")

	// comma,err syntax

	ratings, _ := reader.ReadString('\n')

	input, err := strconv.ParseFloat(strings.TrimSpace(ratings), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating + 1: ", input+1)
	}
}
