package main

import "fmt"

func main() {
	height := 185
	if height >= 180 && height < 200 {
		fmt.Println("In most country you are tall")
	} else if height >= 200 {
		fmt.Println("You are very very tall")
	} else {
		fmt.Println("You of average height")
	}
}
