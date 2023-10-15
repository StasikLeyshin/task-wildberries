/*

Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false


*/

package main

import (
	"fmt"
	"strings"
)

// CheckingUniqueness - проверяет, что все символы в строке уникальные
func CheckingUniqueness(sourceString string) bool {
	// Приводим к нижнему регистру
	sourceString = strings.ToLower(sourceString)

	// Преобразовываем строку в слайс рун
	runeSlice := []rune(sourceString)

	// Мапа для хранения встречающихся символов в строке
	uniquenessMap := make(map[rune]struct{})

	for _, value := range runeSlice {
		_, ok := uniquenessMap[value]
		if ok { // Ели символ value уже существует, то строка не уникальна
			return false
		} else {
			uniquenessMap[value] = struct{}{}
		}
	}

	return true
}

func main() {
	fmt.Println(CheckingUniqueness("abcd"))
	fmt.Println(CheckingUniqueness("abCdefAaf"))
	fmt.Println(CheckingUniqueness("aabcd"))
	fmt.Println(CheckingUniqueness("abcdC"))
	fmt.Println(CheckingUniqueness("AbCdE"))

	/*

		Output:
		true
		false
		false
		false
		false

	*/
}
