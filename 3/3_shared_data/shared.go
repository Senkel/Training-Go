// Go рутины, одновременно работающие с общими данными сами собой не могу синхронизироваться
package main

import (
	"fmt"
)

//Заводим аккаунт
type Account struct {
	balance float64
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {

	a.balance += amount
}

func (a *Account) WithDraw(amount float64) {
	a.balance -= amount
}

func main() {
	acc := Account{}
	//r := rand.Intn(50)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				if j%2 == 1 {
					acc.Deposit(185)

				}
				acc.WithDraw(185)
			}
		}()

	}

	fmt.Println("Balance is", acc)
}
