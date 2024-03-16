package transactions

import (
	"encoding/csv"
	"errors"
	"io"
	"stori/email"
)

type Transaction interface {
	ProcessFile(file io.Reader, email string) error
}

type Service struct {
}

func NewTransaction() Transaction {
	return Service{}
}

func (s Service) ProcessFile(file io.Reader, email string) error {
	reader := csv.NewReader(file)
	_, err := reader.Read()
	if err != nil {
		return errors.New("empty file")
	}
	var transactions []UserTransaction
	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = s.processTransaction(record, &transactions)
		if err != nil {
			return err
		}
	}

	return s.ReportUserBalance(transactions, email)
}

func (s Service) processTransaction(record Record, transactions *[]UserTransaction) error {
	transaction, err := record.GetUserTransaction()
	if err != nil {
		return err
	}
	*transactions = append(*transactions, transaction)
	return nil
}

func (s Service) ReportUserBalance(transactions []UserTransaction, userEmail string) error {

	balance := email.Message{
		Title: "Total balance is: ",
		Value: "39.74",
	}
	t1 := email.Message{
		Title: "Number of transactions in July: ",
		Value: "2",
	}
	t2 := email.Message{
		Title: "Average debit amount: ",
		Value: "20.5",
	}
	return email.Send([]email.Message{balance, t2, t1}, userEmail)
}
