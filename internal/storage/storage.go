package storage

import "PasswordManager/internal/models"

// Storage - интерфейс хранилища
type Storage interface {
	// Login - вход пользователя
	Login(login, password string) error
	// Register - регистрация пользователя
	Register(login, password string) error
	// CheckLogin - проверка существования логина
	CheckLogin(login string) bool
	// SaveEnterData - созранения логина и пароля
	SaveEnterData(login string, data models.EnterData)
	// GetEnterData - получение логина и пароля
	GetEnterData(login, Key string) (models.EnterData, error)
	// SaveTextData - сохранение текста
	SaveTextData(login string, data models.TextData)
	// GetTextData - получение текста
	GetTextData(login, Key string) (models.TextData, error)
	// SaveBankCardData - сохранение банковской карты
	SaveBankCardData(login string, data models.BankCardData)
	// GetBankCardData - получение банковской карты
	GetBankCardData(login, Key string) (models.BankCardData, error)
	// GetAllData - получение всех данных пользователя
	GetAllData(login string) models.AllData
}