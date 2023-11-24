package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for l := j + 1; l <= 9; l++ {
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + j))
				z01.PrintRune(rune('0' + l))
				if i != 7 || j != 8 || l != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')

				}
			}
		}
	}
	z01.PrintRune('\n')
}
