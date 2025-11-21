package structs

import (
	"encoding/json"
	"errors"
	"os"
)

const TRANSACTIONS_FILE = "transactions.json"

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
	fileName := strings.ToLower(TRANSACTIONS_FILE+".json")
	jsonContent, err := json.Marshal(transaction)
	
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
}
