package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func writeBalance(balance float64) {
	os.WriteFile(balanceFile, []byte(fmt.Sprint(balance)), 0644)
}

func readBalance() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 0, errors.New("failed to read balance file (default set to 0)")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("content of balance file is not a number: '" + string(data) + "' (default set to 0)")
	}

	return balance, nil
}

func main() {

	var accountBalance, err = readBalance()

	if err != nil {
		fmt.Println("ERROR: ", err)
		// panic("Cannot conitnue")
	}

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
			writeBalance(accountBalance)
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
			writeBalance(accountBalance)
		} else {
			fmt.Println("Bye!")
			break
		}
	}

}
