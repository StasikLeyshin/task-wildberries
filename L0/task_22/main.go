/*

Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.

*/

package main

import (
	"fmt"
	"math/big"
)

// Addition - сложение
func Addition(a *big.Int, b *big.Int) *big.Int {
	add := new(big.Int).Add(a, b)
	return add
}

// Subtraction - вычитание
func Subtraction(a *big.Int, b *big.Int) *big.Int {
	sub := new(big.Int).Sub(a, b)
	return sub
}

// Multiplication - умножение
func Multiplication(a *big.Int, b *big.Int) *big.Int {
	mul := new(big.Int).Mul(a, b)
	return mul
}

// Division - деление
func Division(a *big.Int, b *big.Int) *big.Int {
	div := new(big.Int).Div(a, b)
	return div
}

func main() {

	// Создаём переменных типа big.Int
	a := big.NewInt(0)
	b := big.NewInt(0)

	// Устанавливаем значения в переменные a и b через SetString, указывая само значение и тип системы счисления
	a.SetString("978697859596978878596586878595678", 10)
	b.SetString("223342515343435254445521543352434", 10)

	fmt.Println(Addition(a, b))       // Сложение
	fmt.Println(Subtraction(a, b))    // Вычитание
	fmt.Println(Multiplication(a, b)) // Умножение
	fmt.Println(Division(a, b))       // Деление

	/*

		Output:
		1202040374940414133042108421948112
		755355344253543624151065335243244
		218584841723625497583617143568382160413706552017021635662143180252
		4

	*/
}
