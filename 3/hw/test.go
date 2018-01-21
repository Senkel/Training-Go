package main

import (
	"fmt"
	"sync"
)

type intCounter struct {
	mu   sync.Mutex
	name string
	age  int
}

/*func (c *intCounter) Add(x int64) {
	c.mu.Lock()
	c.x++
	c.mu.Unlock()
}*/

func (c intCounter) Value(x chan string) {
	d := []string{"g", "o", "l", "a", "n", "g"}
	for i := 0; i < 6; i++ {
		x <- fmt.Sprint(c.name, c.age, " - ", d[i])
	}

}

func main() {
	c := make(chan string)
	counter := intCounter{name: "Vasya ", age: 22}

	go counter.Value(c)

	for i := 0; i < 6; i++ {
		fmt.Println(<-c)
	}

}
