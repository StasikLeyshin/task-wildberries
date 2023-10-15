/*

Реализовать все возможные способы остановки выполнения горутины.

*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Way1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Way1 completed the work")
			return
		}
	}
}

func Way2(stopCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	for {
		select {
		case <-stopCh:
			fmt.Println("Worker Way2 completed the work")
			return
		}
	}
}

func Way3(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.

	time.Sleep(time.Second * 2)
	fmt.Println("Worker Way3 completed the work")
	return
}

func main() {

	// Производим контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Создаём синхронизатор по указателю, для того,
	// чтобы не запутаться и передать по указателю в функцию
	wg := new(sync.WaitGroup)

	// Создаем канал
	intCh := make(chan int)

	// Задаём счётчик горутин
	wg.Add(3)

	go Way1(ctx, wg)   // Завершение через канал контекста
	go Way2(intCh, wg) // Завершение через закрытие канала
	go Way3(wg)        // Самостоятельное завершение

	// Ждём 3 секунды, перед тем, как завершить все горутины
	time.Sleep(time.Second * 3)

	// Уведомляем контекст об его завершении
	cancel()

	// Закрываем канал
	close(intCh)

	// Ждем завершения всех горутин
	wg.Wait()

}
