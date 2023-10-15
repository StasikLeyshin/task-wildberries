/*

Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Writer воркер писателя данных в канал
func Writer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Writer completed the work")
			return
		default:
			ch <- rand.Intn(9999)              // записываем рандомные числа в канал
			time.Sleep(time.Millisecond * 400) // Ожидаем 400 миллисекунд
		}
	}
}

// Reader воркер читателя данных из канала
func Reader(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Reader", "completed the work")
			return
		default:
			fmt.Println("Worker Reader", "Data:", <-ch)
		}
	}
}

func main() {
	var countSecond int

	fmt.Println("Введите количество секунд: ")
	_, err := fmt.Scan(&countSecond)
	if err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
		return
	}

	// Производим контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Создаём синхронизатор по указателю, для того,
	// чтобы не запутаться и передать по указателю в функцию
	wg := new(sync.WaitGroup)

	// Создаем канал
	intCh := make(chan int)

	// Задаём счётчик горутин
	wg.Add(2)

	// Запускаем горутину писателя
	go Writer(ctx, intCh, wg)

	// Запускаем горутину читателя
	go Reader(ctx, intCh, wg)

	// Ждём N секунд, перед завершением программы
	time.Sleep(time.Second * time.Duration(countSecond))

	// Уведомляем контексты об их завершении
	cancel()

	// Закрываем канал
	close(intCh)

	// Ждем завершения всех горутин
	wg.Wait()

}
