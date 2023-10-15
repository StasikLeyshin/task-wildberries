/*

Реализовать бинарный поиск встроенными методами языка.

*/

package main

import "fmt"

// BinarySearch - бинарный поиск
func BinarySearch(arr []int, s int) int {
	var leftPointer = 0             // Установка начальной границы
	var rightPointer = len(arr) - 1 // Установка конечной границы
	for leftPointer <= rightPointer {
		var midPointer = int((leftPointer + rightPointer) / 2) // Средний элемент между начальной и конечной границей
		var midValue = arr[midPointer]

		// Сравнение среднего элемента с искомым значением, при совпадении возвращаем индекс среднего элемента
		if midValue == s {
			return midPointer
		} else {
			// Если средний элемент больше искомого значения, обновим конечную границу (right)
			if midValue < s {
				leftPointer = midPointer + 1
			} else { // Если средний элемент меньше искомого значения, обновим начальную границу (left)
				rightPointer = midPointer - 1
			}
		}
	}
	return -1
}

func main() {
	// Исходный массив
	numbers := []int{13, 21, 24, 54, 63, 85, 93, 112}

	fmt.Println(BinarySearch(numbers, 21))
	fmt.Println(BinarySearch(numbers, 42))
	fmt.Println(BinarySearch(numbers, 112))

	/*

		Output:
		1
		-1
		7

	*/
}
