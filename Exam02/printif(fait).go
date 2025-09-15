package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {
    phrase := Invalid Input
	str1 := []rune(str)
	for i := 0; i <= len(str1); i++ {
		if i >= 3 || i == 3 {
			fmt.Println('G')
			fmt.Println('\n')
		} else {
			fmt.Println("phrase")
			fmt.Println('\n')
			}
	}
	return PrintIf(str)
}
