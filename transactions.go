package main

import (
	"fmt"
	"os"
	"time"

	"example.com/bank/util"
)

const TRANSACTIONS_FILE = "transactions.txt"

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