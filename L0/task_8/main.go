/*

Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

package main

import "fmt"

// SetBit устанавливает i бит в 0 и 1
func SetBit(number int64, i int, valueBit int) int64 {
	if valueBit == 0 {
		number &^= 1 << i
	} else {
		number |= 1 << i
	}
	return number
}

func main() {
	var (
		number   int64 // Исходное число
		i        int   // Номер бита
		valueBit int   // Значение для бита
	)

	fmt.Println("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	fmt.Println("Введите номер бита: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	fmt.Println("Введите значение для бита: [0/1]")
	_, err = fmt.Scan(&valueBit)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	if valueBit < 0 || valueBit > 1 {
		fmt.Printf("Ошибка ввода:")
		return
	}

	fmt.Printf("Исходное число: %d || Исходное число в двоичном виде:%b\n", number, number)
	number = SetBit(number, i, valueBit)
	fmt.Printf("Полученное число: %d || Полученное число в двоичном виде:%b\n", number, number)
}
