// Пакет с реализацией хранилища
package storage

import (
	"PasswordManager/internal/models"
	"errors"
	"sync"
)

// SimpleStorage - простое хранение
type SimpleStorage struct {
	usersData sync.Map
}

// UserCredentials - информация всех пользователей
type UserCredentials struct {
	password    string
	enterData   map[string]models.EnterData
	textData    map[string]models.TextData
	paymentCard map[string]models.PaymentCard
}

var _ Storage = &SimpleStorage{}

// NewSimpleStorage - создание простого хранилища
func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{
		usersData: sync.Map{},
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) Login(login, password string) error {
	var userData any
	var ok bool
	if userData, ok = st.usersData.Load(login); !ok {
		return errors.New("user not exists")
	}
	if userData.(UserCredentials).password != password {
		return errors.New("wrong password")
	}
	return nil
}

// Реализация интерфейса Storage
func (st *SimpleStorage) Register(login, password string) error {
	if _, ok := st.usersData.Load(login); ok {
		return errors.New("user already exists")
	}
	st.usersData.Store(login, UserCredentials{
		password:    password,
		enterData:   make(map[string]models.EnterData),
		textData:    make(map[string]models.TextData),
		paymentCard: make(map[string]models.PaymentCard),
	})
	return nil
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SaveEnterData(login string, data models.EnterData) {
	userData, _ := st.usersData.Load(login)
	userData.(UserCredentials).enterData[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetEnterData(login, key string) (models.EnterData, error) {
	userData, _ := st.usersData.Load(login)
	if enterData, ok := userData.(UserCredentials).enterData[key]; !ok {
		return models.EnterData{}, errors.New("wrong key")
	} else {
		return enterData, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SaveTextData(login string, data models.TextData) {
	userData, _ := st.usersData.Load(login)
	userData.(UserCredentials).textData[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetTextData(login, key string) (models.TextData, error) {
	userData, _ := st.usersData.Load(login)
	if textData, ok := userData.(UserCredentials).textData[key]; !ok {
		return models.TextData{}, errors.New("wrong key")
	} else {
		return textData, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) SavePaymentCard(login string, data models.PaymentCard) {
	userData, _ := st.usersData.Load(login)
	userData.(UserCredentials).paymentCard[data.Key] = data
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetPaymentCard(login, key string) (models.PaymentCard, error) {
	userData, _ := st.usersData.Load(login)
	if paymentPaymentCard, ok := userData.(UserCredentials).paymentCard[key]; !ok {
		return models.PaymentCard{}, errors.New("wrong key")
	} else {
		return paymentPaymentCard, nil
	}
}

// Реализация интерфейса Storage
func (st *SimpleStorage) GetAllData(login string) models.AllData {
	var allData = models.AllData{}
	v, _ := st.usersData.Load(login)
	userData := v.(UserCredentials)
	for _, data := range userData.enterData {
		allData.EnterData = append(allData.EnterData, data)
	}
	for _, data := range userData.paymentCard {
		allData.PaymentCard = append(allData.PaymentCard, data)
	}
	for _, data := range userData.textData {
		allData.TextData = append(allData.TextData, data)
	}
	return allData
}

// Реализация интерфейса Storage
func (st *SimpleStorage) CheckLogin(login string) bool {
	_, ok := st.usersData.Load(login)
	return ok
}
