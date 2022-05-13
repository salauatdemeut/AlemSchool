package main
import "fmt"


func iterativeCalc(index int) int {
	iterativeresult := 0
	for i := 0; i < index + 1; i++ {
		iterativeresult = iterativeresult + i
	}
	return iterativeresult
}
func recursiveCalc(index int) int {
	if index == 1 {
		return 1
	}
	if index > 1 {
		return index + recursiveCalc(index - 1)
	}
	return 0
}

func main() {
	result1 := 0 +1+2+3+4+5+6+7+8+9+10
	fmt.Println(result1)

	result2 := iterativeCalc(13)
	fmt.Printf("The iterative function is: %v\n", result2)

	result3 := recursiveCalc(15)
	fmt.Printf("The recursive function is: %v\n", result3)
}