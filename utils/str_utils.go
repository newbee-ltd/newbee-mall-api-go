package utils

import "fmt"

func SubStrLen(str string, length int) string {
	nameRune := []rune(str)
	fmt.Println("string(nameRune[:4]) = ", string(nameRune[:4]))
	if len(str) > length {
		return string(nameRune[:length-1]) + "..."
	}
	return string(nameRune)
}
