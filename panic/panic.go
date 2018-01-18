package main

import "fmt"

func main() {
	if err := testPanic(); err != nil {
		fmt.Println("Error:", err)
	}
}

func testPanic() (err error) {

	var ok bool
	defer func() {
		if r := recover(); r != nil {
			if err, ok = r.(error); ok {
				fmt.Println("recovered in", r)
			}
			fmt.Println("PANIC Deferred")
		}
	}()

	fmt.Println("start")
	panic(fmt.Errorf("Error occured"))
	return
}
