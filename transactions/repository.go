package transactions

import "github.com/jinzhu/gorm"

type Storage struct {
	Client *gorm.DB
}

func NewStorage(postgres *gorm.DB) Storage {
	return Storage{Client: postgres}
}

func (s Storage) SaveTransactions(transactions UserTransactions) error {
	return s.Client.Transaction(func(tx *gorm.DB) error {
		for _, t := range transactions {
			err := s.Client.Save(t).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s Storage) GetUserTransactions(email string, transactions *UserTransactions) error {
	return s.Client.Where("email = ?", email).Find(transactions).Error
}
