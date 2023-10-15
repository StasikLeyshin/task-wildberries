/*

Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.

*/

package main

import "fmt"

// TypeDefinition - определяет тип переменной по switch type
func TypeDefinition(typeInterface interface{}) {
	switch typeInterface.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan string:
		fmt.Println("chan string")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	// Исходная переменная
	var typeInterface interface{}

	typeInterface = 7482
	TypeDefinition(typeInterface)

	typeInterface = "this is a string"
	TypeDefinition(typeInterface)

	typeInterface = true
	TypeDefinition(typeInterface)

	typeInterface = make(chan string)
	TypeDefinition(typeInterface)
}
