package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - soham@gmail.com"

	file, err := os.Create("./mygofile.txt")

	checkNilErr(err)
	

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	

	fmt.Println("The length of the file is: ", length)
	defer file.Close()
	ReadFile("./mygofile.txt")

}

func ReadFile(path string) {
	databyte, err := ioutil.ReadFile(path)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n",string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}