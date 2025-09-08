package student

import "github.com/01-edu/z01"

func main() {
	var startingASCIINumber int = 122
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(startingASCIINumber - i))
	}
	z01.PrintRune('\n')
}