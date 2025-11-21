package main

import (
	"fmt"
	"time"

	"example.com/bank/structs"
	"example.com/bank/util"
)



func withdrawAmount(balance *float64){
	var withdrawAmount float64
	fmt.Print("Withdraw Amount: ")
	fmt.Scan(&withdrawAmount)
	successOrFailed := "Succeeded"
	if withdrawAmount > *balance{
		fmt.Println("Insufficient Balance")
		successOrFailed = "Failed"
		}else if withdrawAmount <= 0 {
			fmt.Println("Invalid Amount")
			successOrFailed = "Failed"
			}else{
				*balance -= withdrawAmount
				util.WriteFloatToFile(*balance, ACCOUNT_BALANCE_FILE)
	}
	transactionObj, err := structs.NewTransaction("Withdraw: ", withdrawAmount, successOrFailed, time.Now().Format("2006-01-02 15:04:05 Monday"))
	if err != nil {
		fmt.Println("Failed to create transaction object")
		}else{
			transactionObj.Save()
		}
}

func depositAmount(balance *float64){
	var depositAmount float64
	succeedOrFailed := "Succeeded"
	fmt.Print("Deposit Amount: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Invalid Amount")
		succeedOrFailed = "Failed"
		}else {	
			*balance += depositAmount
			util.WriteFloatToFile(*balance, ACCOUNT_BALANCE_FILE)
	}
	transactionObj, err := structs.NewTransaction("Deposit", depositAmount, succeedOrFailed, time.Now().Format("2006-01-02 15:04:05 Monday"))
	if err != nil {
		fmt.Println("Failed to create transaction object")
		}else{
			transactionObj.Save()
		}
}
