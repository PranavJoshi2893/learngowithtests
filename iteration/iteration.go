package iteration

import "strings"

/*
// simple loop

	func Repeat(character string) string {
		var repeated string
		for i := 0; i < 5; i++ {
			repeated += character
		}
		return repeated
	}
*/

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString("a")
	}
	return repeated.String()
}
