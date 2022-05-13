package main

import "github.com/01-edu/z01"

func main() {
	varA := 'A'
	var varB rune = 'B'

	z01.PrintRune(varA)
	z01.PrintRune('\n')
	z01.PrintRune(varB)
	z01.PrintRune('\n')

	// man ascii

	z01.PrintRune(65)
	z01.PrintRune('\n')
	z01.PrintRune(66)
	z01.PrintRune('\n')

	// number use "+"
	z01.PrintRune(65 + 25)
	z01.PrintRune('\n')
}
