package main

import "fmt"

type Structura struct {
	count     int
	firstname string
	Flag      bool
}

func main() {
	s1 := Structura{
		count:     1,
		firstname: "Vasya",
		Flag:      true,
	}
	fmt.Println("Imya - ", s1.firstname)
}
