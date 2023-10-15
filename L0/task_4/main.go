/*

Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
func Reader(ctx context.Context, ch <-chan int, workerId int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда горутина завершается.
	for {
		select {
		case <-ctx.Done(): // Срабатывает, когда переданный в качестве аргумента контекст уведомляет о завершении работы
			fmt.Println("Worker Reader", workerId, "completed the work")
			return
		default:
			fmt.Println("WorkerId:", workerId, "Data:", <-ch)
		}
	}
}

func main() {
	var countWorker int

	fmt.Println("Введите количество воркеров: ")
	_, err := fmt.Scan(&countWorker)
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

	// Задаём счётчик горутин, будет запущено countWorker + 1 горутина на запись
	wg.Add(countWorker + 1)

	fmt.Println("Нажмите Ctrl+C для завершения программы\n")

	// Запускаем горутину писателя
	go Writer(ctx, intCh, wg)

	// Запускаем горутины читателей
	for i := 0; i < countWorker; i++ {
		go Reader(ctx, intCh, i, wg)
	}

	// Создаём канал для обработки сигнала к завершению работы программы через Ctrl + C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)

	// Прослушиваем канал на предмет обнаружения комбинации Ctrl + C
	<-signalChan

	// Уведомляем контексты об их завершении
	cancel()

	// Закрываем канал
	close(intCh)

	// Ждем завершения всех горутин
	wg.Wait()

}
