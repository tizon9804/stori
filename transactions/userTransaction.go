package transactions

import (
	"errors"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserTransaction struct {
	Date            time.Time `gorm:"column:date"`
	TypeTransaction string    `gorm:"column:type_transaction"`
	Transaction     float64   `gorm:"column:transaction"`
	Email           string    `gorm:"column:email"`
}

func (t UserTransaction) TableName() string {
	return "transactions"
}

type UserTransactions []UserTransaction
type Stats struct {
	balance             float64
	transactionPerMonth map[string]int
	avgCredit           float64
	avgDebit            float64
}

func (t *UserTransactions) GetStats() Stats {
	sort.Slice(*t, func(i, j int) bool {
		return (*t)[i].Date.Before((*t)[j].Date)
	})
	balance, credit, debit := float64(0), float64(0), float64(0)
	countCredit, countDebit := float64(0), float64(0)
	monthsCount := map[string]int{}
	for _, t := range *t {
		balance += t.Transaction
		month := t.Date.Month().String()
		monthsCount[month] += 1
		switch t.TypeTransaction {
		case "credit":
			credit += t.Transaction
			countCredit++
		case "debit":
			debit += t.Transaction
			countDebit++
		}
	}
	if countCredit == 0 {
		countCredit = 1
	}
	if countDebit == 0 {
		countCredit = 1
	}
	return Stats{
		balance:             roundMoney(balance),
		transactionPerMonth: monthsCount,
		avgCredit:           roundMoney(credit / countCredit),
		avgDebit:            roundMoney(debit / countDebit),
	}
}

func roundMoney(value float64) float64 {
	return math.Round(value*100) / 100
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
		return UserTransaction{}, errors.New("bad format Transaction value")
	}
	typeTransaction := "debit"
	if strings.Contains(trStr, "+") {
		typeTransaction = "credit"
	}
	return UserTransaction{
		Date:            dateParsed,
		TypeTransaction: typeTransaction,
		Transaction:     value,
	}, nil
}
