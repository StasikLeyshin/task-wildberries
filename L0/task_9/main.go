/*

Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

*/

package main

import (
	"fmt"
	"sync"
)

// ReadArrayWriteChannel - читает слайс и записывает его элементы в канал firstCh
func ReadArrayWriteChannel(wg *sync.WaitGroup, numberList []int, firstCh chan<- int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for index := 0; index < len(numberList); index++ {
		firstCh <- numberList[index]
	}
	close(firstCh)
}

// SquaringWriteChannel - читает канал firstCh и записывает квадраты чисел из канала firstCh в канал secondCh
func SquaringWriteChannel(wg *sync.WaitGroup, firstCh <-chan int, secondCh chan<- int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for number := range firstCh {
		secondCh <- number * number
	}
	close(secondCh)
}

// ReadChannel - читает канал secondCh и выводит сообщения из него
func ReadChannel(wg *sync.WaitGroup, secondCh <-chan int) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for number := range secondCh {
		fmt.Println(number)
	}
}

func main() {
	// Создаём синхронизатор по указателю, для того,
	// чтобы не запутаться и передать по указателю в функцию
	wg := new(sync.WaitGroup)

	// Создаем каналы
	firstCh := make(chan int)
	secondCh := make(chan int)

	// Исходный слайс
	numberList := []int{23, 321, 314, 52, 11}

	// Задаём счётчик горутин
	wg.Add(3)

	// Запускам горутины
	go ReadArrayWriteChannel(wg, numberList, firstCh)
	go SquaringWriteChannel(wg, firstCh, secondCh)
	go ReadChannel(wg, secondCh)

	// Ждем завершения всех горутин
	wg.Wait()

}
