package main
import "fmt"
func main() {
	//sentence := "Hello"
	//sentence :=[]byte("Hello")
	//sentence :=[]rune("Hello")
	//sentence := "中学"
	//sentence :=[]byte("中学")
	sentence :=[]rune("中学")
	counter := 0
	for index, letter := range sentence {
		counter++
		fmt.Printf("index: %v Letter: %c\n", index, letter)
	}
	fmt.Printf("Counter value: %v\n", counter)
	fmt.Println(sentence)

}