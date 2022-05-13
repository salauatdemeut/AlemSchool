package main
import (
	"fmt"
)
func createArrayOfSize (size int) []int {
	//var answer [10]int
	answer := make([]int, size)
	for i := 0; i < size; i++ {
		answer[i] = i +1
	}
	return answer
}
func main() {
	size := 15
	myarray := createArrayOfSize(size)
	fmt.Println(myarray)
}