/*

Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}.
*/

package main

import "fmt"

// GroupingData - объединяет температуры в группы с шагом в 10 градусов
func GroupingData(temperatures []float64, temperaturesGroup map[int][]float64) map[int][]float64 {

	for index := 0; index < len(temperatures); index++ {

		// Получаем ближайший нижний десяток из числа
		key := (int(temperatures[index]) / 10) * 10

		// Добавляем по полученному ключу в мапу
		temperaturesGroup[key] = append(temperaturesGroup[key], temperatures[index])
	}

	return temperaturesGroup
}

// PrintMap - печатает мапу с группировкой температур
func PrintMap(temperaturesGroup map[int][]float64) {
	var k int // Счётчик пройденной мапы, необходим для вывода
	for key, value := range temperaturesGroup {
		fmt.Printf("%d: {", key)

		for index, number := range value {
			fmt.Printf("%.1f", number)
			if index != len(value)-1 {
				fmt.Printf(", ")
			}
		}

		if k != len(temperaturesGroup)-1 {
			fmt.Printf("}, ")
		} else {
			fmt.Printf("}")
		}

		k++
	}
}

func main() {
	// Исходный слайс температур
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Создаём мапу с группировкой температур
	temperaturesGroup := make(map[int][]float64)

	// Группируем
	GroupingData(temperatures, temperaturesGroup)
	// Печатаем
	PrintMap(temperaturesGroup)

	/*

		Output:

		10: {13.0, 19.0, 15.5}, 20: {24.5}, 30: {32.5}, -20: {-25.4, -27.0, -21.0}

	*/
}
