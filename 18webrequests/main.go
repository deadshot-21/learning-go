package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/deadshot-21"

func main() {
	fmt.Println("welcome to web request in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
