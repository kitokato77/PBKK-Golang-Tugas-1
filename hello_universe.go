package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	
	if input == 42 {
		fmt.Println("Universe")
	} else {
		fmt.Println(input)
	}
}
