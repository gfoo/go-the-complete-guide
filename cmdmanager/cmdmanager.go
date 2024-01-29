package cmdmanager

import "fmt"

type CommandManager struct{}

func (cm CommandManager) ReadLines() ([]string, error) {
	fmt.Println("Prices? (finsih with 0)")
	prices := make([]string, 0)
	for {
		var price string
		fmt.Print("> ")
		fmt.Scanln(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	fmt.Println(prices)
	return prices, nil
}

func (cm CommandManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func News() CommandManager {
	return CommandManager{}
}
