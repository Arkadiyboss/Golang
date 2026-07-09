package account

import (
	"encoding/json"
	"fmt"
	"gov1/files"
	"gov1/output"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) string 
}

type Storage struct {
	Accounts    []UserLogin `json:"accounts"`
	UpdatedTime time.Time   `json:"updatedTime"`
}

type StorageWithDb struct {
	Storage Storage `json:"Storage"`
	Db      Db      `json:"-"`
}

func (acc *StorageWithDb) ToBytes() ([]byte, error) {
	acc.Storage.UpdatedTime = time.Now()
	file, error := json.Marshal(acc)
	if error != nil {
		output.PrintError(error)
	}
	return file, error
}

func (newAcc *Storage) AddAccount(acc UserLogin) {
	newAcc.Accounts = append(newAcc.Accounts, acc)
	newAcc.UpdatedTime = time.Now()
}

func NewStorage(db Db) (*StorageWithDb, error) {
	data, err := db.Read()
	if err != nil {
		return &StorageWithDb{
			Storage: Storage{
				Accounts:    []UserLogin{},
				UpdatedTime: time.Now(),
			},
			Db: db,
		}, nil
	}
	var baobab StorageWithDb
	err = json.Unmarshal(data, &baobab)
	if err != nil {
		return nil, err
	}
	return &baobab, nil
}

func (sss *StorageWithDb) FindAccount(delete bool) (*StorageWithDb, error) {
	var inputUrl string
	db := files.NewJsonDb("passwordBase.json")
	fmt.Println("Введите URL")
	fmt.Scan(&inputUrl)
	fmt.Scanln()
	data, err := db.Read()
	if err != nil {
		return nil, err
	}
	var eblan StorageWithDb
	err = json.Unmarshal(data, &eblan)
	if err != nil {
		return nil, err
	}

	var foundAccount StorageWithDb

	for _, account := range eblan.Storage.Accounts {
		hit := strings.Contains(account.Url, inputUrl)
		if hit {
			foundAccount.Storage.Accounts = append(foundAccount.Storage.Accounts, account)
		}
	}
	if len(foundAccount.Storage.Accounts) == 0 {
		fmt.Println("Аккаунты не найдены")
		return sss, nil
	}
	if delete == true {
		return eblan.DeleteAccount(&eblan, inputUrl)
	}
	return &foundAccount, nil
}

func (sss *StorageWithDb) DeleteAccount(eblan *StorageWithDb, inputurl string) (*StorageWithDb, error) {

	var newEblan StorageWithDb

	for index, account := range eblan.Storage.Accounts {
		hit := strings.Contains(account.Url, inputurl)
		if hit != true {
			newEblan.Storage.Accounts = append(newEblan.Storage.Accounts, eblan.Storage.Accounts[index])
		}
	}
	if len(newEblan.Storage.Accounts) > 0 {
		data, err := newEblan.ToBytes()
		if err != nil {
			return nil, err
		}
		db := files.NewJsonDb("passwordBase.json")
		db.Write(data)
	} else {
		fmt.Println("Аккаунт не найден")
	}
	fmt.Println("Данные аккаунта удалены, оставшие аккаунты в списке")
	return &newEblan, nil
}
