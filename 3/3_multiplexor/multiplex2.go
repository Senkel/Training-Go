package main

import (
	"fmt"
)

func fanIn(in, in1 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case x := <-in:
				c <- x
			case x := <-in1:
				c <- x
			}
		}
	}()
	return c
}

func App(s string) <-chan string {
	c := make(chan string)
	go func() {
		c <- fmt.Sprint("beautifull",s)
	}()
	return c
}
func main() {
	// Создаем пару каналов
	z := fanIn(App("Hi"), App("Human"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-z)
	}
}
