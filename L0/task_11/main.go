/*

Реализовать пересечение двух неупорядоченных множеств.

*/

package main

import "fmt"

// Intersection - ищет пересечения двух неупорядоченных множеств
func Intersection(firstList []int, secondList []int) []int {
	intersectionList := make([]int, 0, len(firstList)) // слайс для хранения пересечения двух неупорядоченных множеств

	intersectionMap := make(map[int]struct{}) // map для проверки пересечения множеств

	// Заполняем map числами из первого множества
	for index := 0; index < len(firstList); index++ {
		intersectionMap[firstList[index]] = struct{}{}
	}

	// Проверяем есть ли в map число из первого множества, если есть, то значит есть пересечение
	for index := 0; index < len(secondList); index++ {
		if _, ok := intersectionMap[secondList[index]]; ok {
			intersectionList = append(intersectionList, secondList[index])
		}
	}

	return intersectionList
}

func main() {
	// Первое множество
	firstList := []int{8, 12, 5, 1, 23, 43, 84}

	// Второе множество
	secondList := []int{55, 46, 8, 19, 23, 78, 1}

	fmt.Println(Intersection(firstList, secondList))

}
