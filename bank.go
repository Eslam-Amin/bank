package main

import (
	"fmt"
	"os"
	"time"

	"example.com/bank/util"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"
const TRANSACTIONS_FILE = "transactions.txt"

func main (){
	var accountBalance, err  = util.GetFloatFromFile(ACCOUNT_BALANCE_FILE, 0)
	if err != nil {
		fmt.Println("ERROR\n", err)
	}
	fmt.Println("Welcome to GO Bank!")
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

func withdrawAmount(balance *float64){
	var withdrawAmount float64
	fmt.Print("Withdraw Amount: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount > *balance{
		fmt.Println("Insufficient Balance")
		storeTransactions("Withdrawal: ", withdrawAmount, "Failed")
		}else if withdrawAmount <= 0 {
			fmt.Println("Invalid Amount")
			storeTransactions("Withdrawal: ", withdrawAmount, "Failed")
	}else{
		*balance -= withdrawAmount
		util.WriteFloatToFile(*balance, ACCOUNT_BALANCE_FILE)
		storeTransactions("Withdrawal: ", withdrawAmount, "Succeeded")
	}
}

func depositAmount(balance *float64){
	var depositAmount float64
	fmt.Print("Deposit Amount: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Invalid Amount")
		storeTransactions("Deposit: ", depositAmount, "Failed")
		}else {	
			*balance += depositAmount
			storeTransactions("Deposit: ", depositAmount, "Succeeded")
			util.WriteFloatToFile(*balance, ACCOUNT_BALANCE_FILE)
	}
}


func storeTransactions(transaction string, amount float64, succeedOrFailed string){
	data, _ := os.ReadFile(TRANSACTIONS_FILE)
	textData := string(data)
	currentTime := time.Now().Format("2006-01-02 15:04:05 Monday")
	transactionText := fmt.Sprint(textData ,transaction, amount, " at ", currentTime, " (", succeedOrFailed ,") ", "\n")
	os.WriteFile(TRANSACTIONS_FILE, []byte(transactionText), 0644)	
}

func printTransactions(){
	data, _ := os.ReadFile(TRANSACTIONS_FILE)
	textData := string(data)
	fmt.Println(textData)
}