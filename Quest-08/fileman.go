package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Printf("the mistake is: %v\n", err.Error())
	}
	fmt.Println(file.Stat())

	arr := make([]byte, 6)
	fmt.Println(arr)
	file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
}
