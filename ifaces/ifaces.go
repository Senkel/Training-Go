package main

import "fmt"

type Runer interface {
	Run()
	Alias()
}

type Athlete struct {
	Name string
}

func (a Athlete) Alias() {
	fmt.Println("Alias-Vasya")
}

func (a Athlete) Run() {
	fmt.Println("I can run!,sad", a.Name)
}

func DoRun(r Runer) {
	r.Run()
	r.Alias()
}
func main() {
	athlete := Athlete{"Robert"}

	DoRun(athlete)
	GoRun(athlete)
}

func GoRun(r Runer) {
	r.Alias()
	//b := r.(Athlete) - panic
	//
	if b, ok := r.(Athlete); ok {
		fmt.Println("Klikuxa", b.Name)
	}
}
