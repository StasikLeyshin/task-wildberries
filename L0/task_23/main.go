/*

Удалить i-ый элемент из слайса.

*/

package main

import "fmt"

// RemoveElement - удаляем i-ый элемент из слайса
func RemoveElement(arr []int, index int) []int {
	if len(arr) != 0 && index < len(arr) { // защищаемся от паники
		arr = append(arr[:index], arr[index+1:]...)
	}
	return arr
}

func main() {
	// Индекс элемента, который необходимо удалить
	var index int

	// Исходный слайс
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Введите индекс элемента, который необходимо удалить: ")
	_, err := fmt.Scan(&index)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	fmt.Println("Исходный слайс:", numbers)

	numbers = RemoveElement(numbers, index)

	fmt.Println(numbers)
}
