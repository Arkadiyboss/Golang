package account

import (
	"encoding/json"
	"fmt"
	"gov1/files"
	"os"
	"strings"
	"time"
)

type Storage struct {
	Accounts    []UserLogin `json:"accounts"`
	UpdatedTime time.Time   `json:"updatedTime"`
}

func (acc *Storage) ToBytes() ([]byte, error) {
	acc.UpdatedTime = time.Now()
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

func (sss *Storage) FindAccount(delete bool) (*Storage, error) {
	var inputUrl string
	fmt.Println("Введите URL")
	fmt.Scan(&inputUrl)
	fmt.Scanln()
	data, err := os.ReadFile("passwordBase.json")
	if err != nil {
		return nil, err
	}
	var eblan Storage
	err = json.Unmarshal(data, &eblan)
	if err != nil {
		return nil, err
	}

	var foundAccount Storage

	for _, account := range eblan.Accounts {
		hit := strings.Contains(account.Url, inputUrl)
		if hit {
			foundAccount.Accounts = append(foundAccount.Accounts, account)
		}
	}
	if len(foundAccount.Accounts) == 0 {
		fmt.Println("Аккаунты не найдены")
		return &foundAccount, nil
	}
	if delete {
		return eblan.DeleteAccount(&eblan, inputUrl)
	}
	return &foundAccount, nil
}

func (sss *Storage) DeleteAccount(eblan *Storage, inputurl string) (*Storage, error) {

	var newEblan Storage

	for index, account := range eblan.Accounts {
		hit := strings.Contains(account.Url, inputurl)
		if hit != true {
			newEblan.Accounts = append(newEblan.Accounts, eblan.Accounts[index])
		}
	}
	if newEblan.Accounts != nil {
		data, err := newEblan.ToBytes()
		if err != nil {
			return nil, err
		}
		files.WriteInfo(data, "passwordBase.json")
	}
	fmt.Println("Данные аккаунта удалены, оставшие аккаунты в списке")
	return &newEblan, nil
}
