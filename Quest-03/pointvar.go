package main
import "fmt"
/*
func main() {
	// variables & pointers
	a := 10
	fmt.Println(a)
	fmt.Printf("This is the address of a : %v\n", &a)

	var pointer *int = &a
	fmt.Println(pointer)
}
*/
/*
// bolek func arqilil
func addTen(b int) int {
	b = b+10
	return b
}
func main() {
	a := 10
	a = addTen(a)
	fmt.Println(a)
}
*/

// Endi Pointerdin ozi
func addTen(b *int) {
	*b = *b + 10
	
}
func main() {
	a := 10
	addTen(&a)
	fmt.Println(a)
}