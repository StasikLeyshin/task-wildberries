/*

Реализовать конкурентную запись данных в map.

*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentMap struct {
	mx sync.Mutex // Mьютекс для синхронизации доступа к map
	m  map[string]int
}

// NewParallelMap возвращает новый ConcurrentMap
func NewParallelMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]int),
	}
}

// Get Метод для получения значения по ключу из map
func (c *ConcurrentMap) Get(key string) (int, bool) {
	c.mx.Lock()         // Захватываем мьютекс получением значения
	defer c.mx.Unlock() // Освобождаем мьютекс после завершения функции
	value, ok := c.m[key]
	return value, ok
}

// Set метод для установки значения по ключу в map
func (c *ConcurrentMap) Set(key string, value int) {
	c.mx.Lock()         // Захватываем мьютекс перед записью
	defer c.mx.Unlock() // Освобождаем мьютекс после завершения функции
	c.m[key] = value
}

// Writer воркер писателя данных в конкурентную мапу
func Writer(ctx context.Context, wg *sync.WaitGroup, ID int, concurrentMap *ConcurrentMap) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Writer", ID, "completed the work")
			return
		default:
			number := rand.Intn(9999)
			concurrentMap.Set(fmt.Sprintf("Worker ID %d", ID), number) // записываем данные в map
			time.Sleep(time.Second)                                    // Ожидаем 400 миллисекунд
		}
	}
}

func main() {
	concurrentMap := NewParallelMap()

	// Производим контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Создаём синхронизатор по указателю, для того,
	// чтобы не запутаться и передать по указателю в функцию
	wg := new(sync.WaitGroup)

	// Задаём счётчик горутин
	wg.Add(10)

	// Запускаем писателей в конкурентную map
	for i := 0; i < 10; i++ {
		go Writer(ctx, wg, i, concurrentMap)
	}

	time.Sleep(time.Second * 2)

	// Выводим всю map
	for key, value := range concurrentMap.m {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Уведомляем контекст об его завершении
	cancel()

	// Ждем завершения всех горутин
	wg.Wait()
}
