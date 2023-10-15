/*

Реализовать паттерн «адаптер» на любом примере.

*/

package main

import "fmt"

// InstructionManual - инструкция
type InstructionManual interface {
	Read()
}

// English - инструкция на английском
type English struct {
}

// NewEnglish - конструктор для English
func NewEnglish() *English {
	return &English{}
}

// Read - метод чтения инструкции на английском
func (e *English) Read() {
	fmt.Println("Reading the instructions in English")
}

// Russian - инструкция на русском
type Russian struct {
}

// Read - метод чтения инструкции на русском
func (a *Russian) Read() {
	fmt.Println("Чтение инструкции на русском")

}

// RussianAdapter - Адаптер для русской версии инструкции
type RussianAdapter struct {
	RussianVersion *Russian
}

// NewRussianAdapter - Конструктор для адаптера
func NewRussianAdapter() InstructionManual {
	return &RussianAdapter{}
}

// Read - метод перевода инструкции с английского на русский
func (r *RussianAdapter) Read() {
	fmt.Println("Адаптер переводит английский на русский")
	r.RussianVersion.Read()
}

func main() {
	en := NewEnglish()
	rus := NewRussianAdapter()

	en.Read()
	rus.Read()
}
