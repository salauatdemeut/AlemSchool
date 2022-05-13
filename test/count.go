package main
import (
	"github.com/01-edu/z01"
	 "os"

)
func main() {
	 args1 := os.Args[1:]
	result := ""
	n := len(args1)
	if n == 0{
		z01.PrintRune('0')
		return
	}
	for n > 0 {
		result += string('0' + n % 10)
		n /= 10
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(rune(result[i]))
		
	}
	z01.PrintRune('\n')
}