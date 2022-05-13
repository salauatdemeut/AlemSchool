package main
import (
	"strconv"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	a, _ := strconv.Atoi(arg)
	var b int
	if a > 0 {
		if len(a) == 1 {
			if b = a/b {
				print("true")
			} else {
				print("false")
			}
		}
	}
}