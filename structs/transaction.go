package structs

import (
	"encoding/json"
	"errors"
	"os"
)

const TRANSACTIONS_FILE = "transactions.json"
var Transactions []Transaction

type Transaction struct {
	TypeOfTransaction string `json:"typeOfTransaction"`
	Amount float64 `json:"amount"`
	SucceedOrFailed string `json:"succeedOrFailed"`
	CreatedAt string `json:"createdAt"`
}


func NewTransaction(typeOfTransaction string, amount float64, succeedOrFailed string, createdAt string) (*Transaction, error){
	if typeOfTransaction == "" || amount <= 0 || succeedOrFailed == "" || createdAt == "" {
		return nil, errors.New("type Of Transaction, amount, succeed Or Failed, and createdAt are required")
	}

	return &Transaction{
		TypeOfTransaction: typeOfTransaction,
		Amount: amount,
		SucceedOrFailed: succeedOrFailed,
		CreatedAt: createdAt,
	}, nil
}

func (transaction Transaction) Save() error{

	listOfTransactions, err := LoadTransactions()
	if err != nil {
		return err
	}
	listOfTransactions = append(listOfTransactions, transaction)
	Transactions = listOfTransactions
	jsonContent, err := json.Marshal(listOfTransactions)
	if err != nil {
		return err
	}
	return os.WriteFile(TRANSACTIONS_FILE, jsonContent, 0644)
}


func LoadTransactions() ([]Transaction, error) {  
	var list []Transaction
		if len(Transactions) > 0 {
			return Transactions, nil
		}
    data, err := os.ReadFile(TRANSACTIONS_FILE)
    if err != nil {
        return nil, err
    }
		if len(data) ==0 {
			Transactions = list
			return list, nil
		}
		err = json.Unmarshal(data, &list);
    if  err != nil {
			return nil, err
    }
		
		Transactions = list
    return list, nil
}