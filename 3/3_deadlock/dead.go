package main

import (
	"fmt"
	"time"
)

type Word struct {
	count int
}

func main() {
	c := make(chan *Word)

	go game("Тяп", c)
	go game("Ляп", c)
	c <- new(Word) //запустить слово в игру , иначе deadlock
	time.Sleep(time.Second)
	<-c

}

func game(s string, word chan *Word) {
	for {
		words := <-word
		words.count++
		fmt.Println(s, words.count)
		time.Sleep(100 * time.Millisecond)
		word <- words
	}

}
