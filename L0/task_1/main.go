/*

Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

*/

package main

import "fmt"

type Human struct {
	age    int
	gender string
}

// GetAge получаем возраст человека
func (h *Human) GetAge() int {
	return h.age
}

// SetAge устанавливаем возраст человека
func (h *Human) SetAge(age int) {
	h.age = age
}

// GetGender получаем пол человека
func (h *Human) GetGender() string {
	return h.gender
}

// SetGender устанавливаем пол человека
func (h *Human) SetGender(gender string) {
	h.gender = gender
}

type Action struct {
	Human // Встраиваем тип Human в структуру Action
}

// LiveYear Прибавка года к возрасту человека
func (a *Action) LiveYear() {
	a.Human.SetAge(a.Human.GetAge() + 1)
	fmt.Println("Human lived for one year. He turned:", a.Human.GetAge())
}

func main() {
	action := Action{Human{12, "male"}}
	fmt.Printf("Parameters Human: %d old, %s gender\n", action.GetAge(), action.GetGender())
	action.LiveYear()

	/*

		Output:
		Parameters Human: 12 old, male gender
		Human lived for one year. He turned: 13

	*/
}
