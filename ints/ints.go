package main

import (
	"fmt"
)

type MyInt int

type MyStruct struct {
	Num  int
	Name string
}

func (m MyInt) showYourSelf() {
	fmt.Println("AGA:", m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

func main() {
	i := MyInt(2)
	i.showYourSelf()
	z := MyInt(23)
	i.add(z)
	fmt.Println(i)
}
