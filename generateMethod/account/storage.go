package account

import (
	"encoding/json"
	"os"
	"time"
)

type Storage struct {
	Accounts    []UserLogin `json:"accounts"`
	UpdatedTime time.Time   `json:"updatedTime"`
}

func (acc *Storage) ToBytes() ([]byte, error) {
	file, error := json.Marshal(acc)
	if error != nil {
		return nil, error
	}
	return file, error
}

func (newAcc *Storage) AddAccount(acc UserLogin) {
	newAcc.Accounts = append(newAcc.Accounts, acc)
	newAcc.UpdatedTime = time.Now()
}

func NewStorage() (*Storage, error) {
	data, err := os.ReadFile("passwordBase.json")
	if err != nil {
		return &Storage{
			Accounts:    []UserLogin{},
			UpdatedTime: time.Now(),
		}, nil
	}
	var baobab Storage
	err = json.Unmarshal(data, &baobab)
	if err != nil {
		return nil, err
	}
	return &baobab, nil
}
