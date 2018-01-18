package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Surname string
	Age     int
}

//
type Сharacter struct {
	Behavior string
	Iq       int
}

type CompletePerson struct {
	Person
	Сharacter
	Appearance string
}

func (p Person) GetName() string {
	fmt.Println(p.Name)
	return p.Name
}

func (cp CompletePerson) GetCompletePerson() Person {
	return cp.Person
}

func (cp CompletePerson) GetIdealCharacter() Сharacter {
	return cp.Сharacter
}

func main() {
	per := Person{"Mantsevich", "Artur", 15}
	fmt.Println("Name of person", per.GetName())
	ip := CompletePerson{Person: Person{"Diana", "Shuleiko", 18}, Сharacter: Сharacter{"Sweetheart", 200}, Appearance: "Beautiful"}
	fmt.Print("Ideal person", ip.GetCompletePerson())
	fmt.Println("Ideal character", ip.GetIdealCharacter())

}
