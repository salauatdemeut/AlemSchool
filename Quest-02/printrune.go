package main

import "github.com/01-edu/z01"

func main() {
	
	// Just Print Rune
	var aRune rune = 'A'
	z01.PrintRune(aRune)
	z01.PrintRune('\n')


	// String to Rune
	var aRunes string = "Hello"
	z01.PrintRune(rune(aRunes[0]))
	z01.PrintRune(rune(aRunes[1]))
	z01.PrintRune(rune(aRunes[2]))
	z01.PrintRune(rune(aRunes[3]))
	z01.PrintRune(rune(aRunes[4]))
	z01.PrintRune('\n')
}
