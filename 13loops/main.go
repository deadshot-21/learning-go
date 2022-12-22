package main

import "fmt"

func main() {

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	fmt.Println(days)

	for d:=0; d<len(days); d++ {
		fmt.Println(days[d])
	}

	for i:= range days {
		fmt.Println(days[i])
	}

	for _, day := range days{
		fmt.Println(day)
	}

	rougueValue := 1

	for rougueValue < 10{
		if rougueValue == 3{
			rougueValue++
			continue
		}

		if rougueValue ==4 {
			goto lco
		}

		if rougueValue ==5 {
			break
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println(("Jumped here"))
}