package main

import (
	"fmt"
)

func boring(s string) <-chan string {

	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", s, i)
		}
	}()

	return c
}

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say:  \n", <-c)
	}
}
