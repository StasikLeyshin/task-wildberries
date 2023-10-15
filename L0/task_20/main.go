/*

Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

*/

package main

import (
	"fmt"
	"strings"
)

// ReverseString - переворачивает слова входной строки
func ReverseString(sourceString string) string {

	// Разбивает строку на слайс строк
	runeSlice := strings.Split(sourceString, " ")

	for i := 0; i < len(runeSlice)/2; i++ {
		runeSlice[i], runeSlice[(len(runeSlice)-1)-i] = runeSlice[(len(runeSlice)-1)-i], runeSlice[i]
	}

	// Преобразовываем слайс строк в строку
	return strings.Join(runeSlice, " ")
}

func main() {
	fmt.Println(ReverseString("snow dog sun"))
	fmt.Println(ReverseString("замок раз два три"))
	fmt.Println(ReverseString("красное лицо смотрело на меня"))

	/*

		Output:
		sun dog snow
		три два раз замок
		меня на смотрело лицо красноел

	*/
}
