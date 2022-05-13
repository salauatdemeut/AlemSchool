package main
import ("github.com/01-edu/z01"
"os")

func main () {
	args := os.Args[1:]
	count := '0'

	for i := 0; i < len(args); i++ {
		count++
	}
	z01.PrintRune(count)
	z01.PrintRune('\n')
}