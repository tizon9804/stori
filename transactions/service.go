package transactions

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"stori/email"
)

type Transaction interface {
	ProcessFile(file io.Reader, email string) error
}

type Service struct {
	storage Storage
}

func NewTransaction(storage Storage) Transaction {
	return Service{
		storage: storage,
	}
}

func (s Service) ProcessFile(file io.Reader, email string) error {
	reader := csv.NewReader(file)
	_, err := reader.Read()
	if err != nil {
		return errors.New("empty file")
	}
	var transactions UserTransactions
	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = s.processTransaction(record, email, &transactions)
		if err != nil {
			return err
		}
	}
	err = s.storage.SaveTransactions(transactions)
	if err != nil {
		return err
	}
	var transactionsDB UserTransactions
	err = s.storage.GetUserTransactions(email, &transactionsDB)
	if err != nil {
		return err
	}
	s.ReportUserBalance(transactionsDB, email)
	return nil
}

func (s Service) processTransaction(record Record, email string, transactions *UserTransactions) error {
	transaction, err := record.GetUserTransaction()
	if err != nil {
		return err
	}
	transaction.Email = email
	*transactions = append(*transactions, transaction)
	return nil
}

func (s Service) ReportUserBalance(transactions UserTransactions, userEmail string) {
	stats := transactions.GetStats()
	balance := email.Message{
		Title: "Total balance is: ",
		Value: fmt.Sprintf("%v", stats.balance),
	}
	var transactionsMonth []email.Message
	for month, count := range stats.transactionPerMonth {
		tm := email.Message{
			Title: fmt.Sprintf("Number of transactions in %s:", month),
			Value: fmt.Sprintf("%v", count),
		}
		transactionsMonth = append(transactionsMonth, tm)
	}
	avgDebit := email.Message{
		Title: "Average debit amount: ",
		Value: fmt.Sprintf("%v", stats.avgDebit),
	}
	avgCredit := email.Message{
		Title: "Average credit amount:",
		Value: fmt.Sprintf("%v", stats.avgCredit),
	}
	messages := []email.Message{balance, avgCredit, avgDebit}
	messages = append(messages, transactionsMonth...)

	go func() {
		err := email.Send(messages, userEmail)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Email Sent")
	}()
}
