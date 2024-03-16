package transactions

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type UserTransaction struct {
	date            time.Time
	typeTransaction string
	transaction     float64
}

type Record []string

func (t Record) GetUserTransaction() (UserTransaction, error) {
	dateStr := t[0]
	trStr := t[1]
	dateParsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return UserTransaction{}, err
	}
	value, err := strconv.ParseFloat(trStr, 64)
	if err != nil {
		return UserTransaction{}, errors.New("bad format transaction value")
	}
	typeTransaction := "debit"
	if strings.Contains(trStr, "+") {
		typeTransaction = "credit"
	}
	return UserTransaction{
		date:            dateParsed,
		typeTransaction: typeTransaction,
		transaction:     value,
	}, nil
}
