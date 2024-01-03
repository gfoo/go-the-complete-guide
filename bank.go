package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"org.gfoo/bank/ops"
)

const balanceFile = "balance.txt"

func main() {

	var accountBalance, err = ops.ReadFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR: ", err)
		// panic("Cannot conitnue")
	}

	fmt.Println("Welcome to GO bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		printChoice()

		var choice int
		fmt.Print("Choice? ")
		fmt.Scan(&choice)

		wantCheckBalance := choice == 1
		wantDeposit := choice == 2
		wantWitdraw := choice == 3

		if wantCheckBalance {
			fmt.Println("Balance : ", accountBalance)
		} else if wantDeposit {
			fmt.Print("Desposit? ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Deposit must be > 0")
				continue
			}
			accountBalance += deposit
			fmt.Println("New Balance : ", accountBalance)
			ops.WriteFloatToFile(balanceFile, accountBalance)
		} else if wantWitdraw {
			fmt.Print("Withdraw? ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Withdraw must be > 0")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Withdraw must be <= ", accountBalance)
				continue
			}
			accountBalance -= withdraw
			fmt.Println("New Balance : ", accountBalance)
			ops.WriteFloatToFile(balanceFile, accountBalance)
		} else {
			fmt.Println("Bye!")
			break
		}
	}

}
