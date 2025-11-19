package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func main (){
	var accountBalance, err  = readBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR\n", err)
	}
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("What do you want to do?")
	for{
		actionMenu()
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
			fmt.Println("Thanks for choosing our bank")
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Choice")
		}
	}
}



func actionMenu(){
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func withdrawAmount(balance *float64){
	var withdrawAmount float64
	fmt.Print("Withdraw Amount: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > *balance{
		fmt.Println("Insufficient Balance")
		}else if withdrawAmount <= 0 {
			fmt.Println("Invalid Amount")
	}else{
		*balance -= withdrawAmount
		writeBalanceToFile(*balance)
	}
}

func depositAmount(balance *float64){
	var depositAmount float64
	fmt.Print("Deposit Amount: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Invalid Amount")
		}else {	
			*balance += depositAmount
			writeBalanceToFile(*balance)
	}
}

func writeBalanceToFile(balance float64){
	textBalance := fmt.Sprint(balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(textBalance), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err :=os.ReadFile(ACCOUNT_BALANCE_FILE)
	if err != nil{
		return 0, errors.New("failed to find balance file")
	}
	textBalance := string(data)
	currentBalance, err := strconv.ParseFloat(textBalance, 64)
	if err != nil{
		return 0, errors.New("failed to parse stored balance value")
	}
	return currentBalance, nil
}

