package main
import (
	"fmt"
)
func createArrayOfSize (size int) []int {
	//var answer [10]int
	var answer[]int
	for i := 0; i < size; i++ {
		answer = append(answer, i +1)
	 // answer[i] = i +1
	}
	return answer
}
func main() {
	size := 15
	myarray := createArrayOfSize(size)
	fmt.Println(myarray)
}