package main

import (
	"fmt"
)

func main() {
	TopUpAmount(100000, 2000)
	PayCredit(3000000, 263.00)
}
func TopUpAmount(balance int, TopUp int) {
	month := 0
	for {
		month++
		if month > 12 {
			month = 1
		}
		fmt.Println("Месяц: ", month)
		balance += TopUp
		fmt.Println("Баланс: ", balance)
		if balance > 500000 {
			fmt.Println("Достигнут предел")
			break
		}
	}
}

 func PayCredit(Credit int, PayMonth int) {
 	month := 0
	rate := 10
 	for {
 		month++
 		if month > 12 {
 			month = 1
 		}
		fmt.Println("Month: ", month)
		Credit -= PayMonth * rate/100
		fmt.Println("Credit: ", Credit)
		if Credit < 500000 {
			fmt.Println("A little bit more")
			break
		}
 	}
 }
