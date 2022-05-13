package main
import "fmt"
func main() {
	slice := []string{
		"Word1",
		"Word2",
		"Word3",
		"Word4",
		"Word5",
		"Word6",
	}

	for index, word := range slice {
		fmt.Printf("The index is: %v, The element is: %v\n", index, word)
	}
	slice2 := []string{
		"Word7",
		"Word8",
		"Word9",
		"Word10",
		"Word11",
		"Word12",
	}

	for _, word2 := range slice2 {
		fmt.Printf("The element is: %v\n", word2)
	}
}