/*

К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


Данный фрагмент кода может привести к выделению большого объема памяти для строки v,
что может вызвать проблемы с использованием памяти и привести к сбою программы.

Поэтому необходимо воспользоваться буфером, который эффективно создаёт и
изменяет строки без выделения большого объема памяти.

*/

package main

import (
	"bytes"
)

var justString string

func createHugeString(stringLen int) string {
	var buf bytes.Buffer
	for i := 0; i < stringLen; i++ {
		buf.WriteString("Test")
	}
	return buf.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
