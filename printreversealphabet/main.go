package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := len(alphabet) - 1; i >= 0; i-- {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}
