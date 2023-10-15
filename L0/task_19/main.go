/*

Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

*/

package main

import "fmt"

// ReverseString - переворачивает символы входной строки
func ReverseString(sourceString string) string {
	// Преобразовываем строку в слайс рун
	runeSlice := []rune(sourceString)

	for i := 0; i < len(runeSlice)/2; i++ {
		runeSlice[i], runeSlice[(len(runeSlice)-1)-i] = runeSlice[(len(runeSlice)-1)-i], runeSlice[i]
	}

	// Преобразовываем слайс рун в строку
	return string(runeSlice)
}

func main() {
	fmt.Println(ReverseString("главрыба"))
	fmt.Println(ReverseString("замок"))
	fmt.Println(ReverseString("лицо"))

	/*

		Output:
		абырвалг
		комаз
		оцил

	*/
}
