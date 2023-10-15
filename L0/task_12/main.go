/*

Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.


*/

package main

import "fmt"

// CreateSet - создаёт собственные множества
func CreateSet(stringList []string) []string {
	stringMap := make(map[string]struct{}) // map для хранения множеств

	setList := make([]string, 0, len(stringList)) // для вывода всех множеств

	// Заполняем map
	for _, value := range stringList {
		stringMap[value] = struct{}{}
	}

	// Добавляем в слайс множества
	for key, _ := range stringMap {
		setList = append(setList, key)
	}
	return setList
}

func main() {
	// исходный слайс
	stringList := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(CreateSet(stringList))
}
