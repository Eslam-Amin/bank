package main

import (
	"fmt"

	"example.com/bank/util"
	"github.com/Pallinder/go-randomdata"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func main (){
	var accountBalance, err  = util.GetFloatFromFile(ACCOUNT_BALANCE_FILE, 0)
	if err != nil {
		fmt.Println("ERROR\n", err)
	}
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7 at: ",randomdata.PhoneNumber())
	fmt.Println("Your account Name is: ", randomdata.SillyName())
	fmt.Println("What do you want to do?")
	for{
		optionsMenu()
		fmt.Print("Your Choice: ")
		var choice int 
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Balance: ", accountBalance)
			fmt.Println("########################")
		case 2:
			depositAmount(&accountBalance)
			fmt.Println("Balance: ", accountBalance)
			fmt.Println("########################")
		case 3:
			withdrawAmount(&accountBalance)
			fmt.Println("Balance: ", accountBalance)
			fmt.Println("########################")
		case 4:
			printTransactions()
			fmt.Println("########################")
		case 5:
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}


