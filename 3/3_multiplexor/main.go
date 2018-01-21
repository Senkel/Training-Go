package main

import (
	"fmt"
)

func funIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {

		for {
			select {
			case x := <-input1:
				c <- x
			case x := <-input2:
				c <- x
			}
		}
	}()

	return c
}

func main() {
	c := funIn(Boring("hei"), happy("human"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c, "!")
	}
}

func Boring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- fmt.Sprintf(s)
		}
	}()
	return c
}
func happy(s string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprint("beautiful ", s)
		}
	}()
	return c
}
