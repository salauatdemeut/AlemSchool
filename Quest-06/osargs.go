package main
import (
	"os"
	"fmt"
)
func main() {
	arguments := os.Args
	fmt.Printf("Number of arguments: %v\n", len(arguments))

	if len(arguments) >=3 {
		fmt.Printf("arguments #2 is: %v\n", os.Args[3])
	}
}