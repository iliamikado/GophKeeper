// Пакет с реализацией хранилища
package storage

import (
	"PasswordManager/internal/models"
	"errors"
)

// SimpleStorage - простое хранение
type SimpleStorage struct {
	usersData map[string] UserData
}

// UserData - информация всех пользователей
type UserData struct {
	password string
	enterData map[string]models.EnterData
	textData map[string]models.TextData
	bankCardData map[string]models.BankCardData
}

var _ Storage = &SimpleStorage{}

// NewSimpleStorage - создание простого хранилища
func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		usersData: make(map[string]UserData),
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) Login(login, password string) error {
	var userData UserData
	var ok bool
	if userData, ok = st.usersData[login]; !ok {
		return errors.New("user not exists")
	}
	if userData.password != password {
		return errors.New("wrong password")
	}
	return nil
}

// Реализация интерфейса Storage
func (st *SimpleStorage) Register(login, password string) error {
	if _, ok := st.usersData[login]; ok {
		return errors.New("user already exists")
	}
	st.usersData[login] = UserData{
		password: password,
		enterData: make(map[string]models.EnterData),
		textData: make(map[string]models.TextData),
		bankCardData: make(map[string]models.BankCardData),
	}
	return nil
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SaveEnterData(login string, data models.EnterData) {
	st.usersData[login].enterData[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetEnterData(login, key string) (models.EnterData, error) {
	if enterData, ok := st.usersData[login].enterData[key]; !ok {
		return models.EnterData{}, errors.New("wrong key")
	} else {
		return enterData, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SaveTextData(login string, data models.TextData) {
	st.usersData[login].textData[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetTextData(login, key string) (models.TextData, error) {
	if textData, ok := st.usersData[login].textData[key]; !ok {
		return models.TextData{}, errors.New("wrong key")
	} else {
		return textData, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SaveBankCardData(login string, data models.BankCardData) {
	st.usersData[login].bankCardData[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetBankCardData(login, key string) (models.BankCardData, error) {
	if bankCardData, ok := st.usersData[login].bankCardData[key]; !ok {
		return models.BankCardData{}, errors.New("wrong key")
	} else {
		return bankCardData, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetAllData(login string) models.AllData {
	var allData = models.AllData{}
	for _, data := range st.usersData[login].enterData {
		allData.EnterData = append(allData.EnterData, data)
	}
	for _, data := range st.usersData[login].bankCardData {
		allData.BankCardData = append(allData.BankCardData, data)
	}
	for _, data := range st.usersData[login].textData {
		allData.TextData = append(allData.TextData, data)
	}
	return allData
}

// Реализация интерфейса Storage
func (st *SimpleStorage) CheckLogin(login string) bool {
	_, ok := st.usersData[login]
	return ok;
}