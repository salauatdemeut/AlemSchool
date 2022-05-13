package main

import (
	"github.com/01-edu/z01"
)

func main() {
	mode := 0 // 0 = small; 1 = big
	for i := 'z'; i >= 'a'; i-- {
		if mode == 0 {
			z01.PrintRune(i)
			mode = 1
		} else if mode == 1 {
			z01.PrintRune(i - 32)
			mode = 0
		}
	}
	z01.PrintRune('\n')
}
