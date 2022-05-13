package main

import "fmt"

func addTen(nbr int) int {
	return nbr + 10
}

func addTwenty(nbr int) int {
	return nbr + 20
}

func addThirty(nbr int) int {
	return nbr + 30
}

func applyFunction(f func(int) int, answer int) int {
	answer = f(answer)
	return answer
} 

func main() {
	result := 0

	arrayOfFunction := []func(int) int{addTen, addTwenty, addThirty}
	result = applyFunction(arrayOfFunction[0], result)
	// result = applyFunction(addTen, result)
	fmt.Println(result)
}
