/*

Поменять местами два числа без создания временной переменной.

*/

package main

import "fmt"

// Way1 - меняем местами две переменные с помощью двойного присваивания
func Way1(a int, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

// Way2 - меняем местами две переменные с помощью сложения и вычитания
func Way2(a int, b int) {
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}

// Way3 - меняем местами две переменные с помощью умножения и деления
func Way3(a int, b int) {
	a = a * b
	b = a / b
	a = a / b
	fmt.Println(a, b)
}

func main() {
	a := 5
	b := 9

	fmt.Println(a, b)

	Way1(a, b)
	Way2(a, b)
	Way3(a, b)
}
