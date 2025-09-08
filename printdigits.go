package main

import "github.com/01-edu/z01"

func main() {
	var startingASCIINumber int = 48
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(startingASCIINumber + i))
	}
	z01.PrintRune('\n')
}