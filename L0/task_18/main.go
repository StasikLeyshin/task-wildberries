/*

Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.

*/

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	num int
	mx  sync.Mutex
}

// NewCounter - конструктор для структуры счётчика Counter
func NewCounter() *Counter {
	return &Counter{
		num: 0,
	}
}

// Inc - инкрементация счётчика в конкурентной среде
func (c *Counter) Inc() {
	c.mx.Lock()         // Захватываем мьютекс получением значения
	defer c.mx.Unlock() // Освобождаем мьютекс после завершения функции
	c.num++             // Увеличиваем значение счётчика на 1
}

// Value - получение значения счётчика
func (c *Counter) Value() int {
	return c.num
}

// Increment - запускает инкремент счётчика
func Increment(wg *sync.WaitGroup, counter *Counter) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	counter.Inc()
}

func main() {
	counter := NewCounter()

	// Количество запущенных горутин
	countGoroutine := 45

	// Создаём синхронизатор
	wg := new(sync.WaitGroup)

	// Задаём счётчик горутин
	wg.Add(countGoroutine)

	for i := 0; i < countGoroutine; i++ {
		go Increment(wg, counter)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Печатаем значение счётчика
	fmt.Println(counter.Value())
}
