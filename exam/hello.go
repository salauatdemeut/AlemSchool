package main 
import "github.com/01-edu/z01"
func main() {
	s := "Hello World"
	r :=[]rune(s)
	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}