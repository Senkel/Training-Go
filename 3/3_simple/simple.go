// Запуск go рутин
// именованные функции или анонимные
package main

import (
	"fmt"
)

func main() {
	fmt.Println("старт")
	// можно запустить функцию
	go process(0)
	// можно запустить анонимную функцию
	go func() {
		fmt.Println("Анонимный запуск")
	}()

	// Можем запустить много горутин
	for i := 0; i < 10; i++ {
		go process(i)
	}

	n := Name("Vladislav")
	go n.run()
	// Нужно дождаться заверешния выполнения
	fmt.Scanln()

}
func (n Name) run() {
	println("koko")
}

type Name string

func process(i int) {
	fmt.Println("обработка: ", i)
}
