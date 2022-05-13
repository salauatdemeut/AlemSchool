package main
import (
	"fmt"
)
type student struct {
	name string
	age int
}

func main() {
	chris := student{}
	fmt.Println(chris)

	chris.name = "chris"
	chris.age = 30
	fmt.Println(chris)
}