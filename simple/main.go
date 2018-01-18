package main

import "fmt"

var st = "not ready"

func main() {
	stuff := []int{1, 2, 10}

	fmt.Println("res", running(stuff...))
	fmt.Println("hi!")
	fmt.Println("The stuff is", st) //ready

}

func init() {
	st = "ready"
}

func running(stuf ...int) (res int) {
	for i := range stuf {
		res += stuf[i]
	}
	return
}
