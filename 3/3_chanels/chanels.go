// Один из механизмов синхронизации - каналы
// Каналы, это объект через который можно обеспечить взаимодействие нескольких горутин
// В принимающей (или возвращающей) канал функции, можно указать направление работы с каналом
// Только для чтения - "<-chan" или только для записи "chan<-"
package main

import (
	"fmt"
	"math/rand"
)

func greet(c chan<- string) {
	// Запускаем бесконечный цикл
	for {
		// и пишем в канал пару строк
		// Подпрограмма будет заблокирована до того, как кто-то захочет прочитать из канала
		c <- fmt.Sprintf("Магистр")
		c <- fmt.Sprintf("Ёда")
	}
}

func process(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}

func main() {
	// Создаем канал
	c := make(chan string)

	// стартуем пишущую горутину
	r := rand.Intn(50)
	go greet(c)
	go Out(c)
	for i := 0; i < r; i++ {
		// Читаем пару строк из канала
		fmt.Println(<-c, ",", <-c, "->", <-c)
	}

	stuff := make(chan int, 50)
	for i := 0; i < 55; i = i + 3 {
		stuff <- i
	}
	close(stuff)
	fmt.Println("Res", process(stuff))
}

func Out(x chan<- string) {
	for {
		x <- fmt.Sprintf("Out")
	}

}
