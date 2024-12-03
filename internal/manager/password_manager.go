// Пакет с основной логикой
package manager

import (
	"PasswordManager/internal/models"
	"PasswordManager/internal/storage"
	"math/rand"
)

// IPasswordManager - струкура основного приложения
type IPasswordManager struct {
	st storage.Storage
}

// PasswordManager - основнай экзэмпляр реализующий логику
var PasswordManager *IPasswordManager

// Initialize - инициализация PasswordManager
func Initialize() {
	PasswordManager = &IPasswordManager{
		st: storage.NewSimpleStorage(),
	}
}

// Register - регистрация пользователя
func (pm *IPasswordManager) Register(login, password string) (string, error) {
	err := pm.st.Register(login, password)
	if err != nil {
		return "", err
	}
	return login, nil
}

// Login - вход пользователя
func (pm *IPasswordManager) Login(login, password string) (string, error) {
	err := pm.st.Login(login, password)
	if err != nil {
		return "", err
	}
	return login, nil
}

// SaveEnterData - сохранение логина и пароля
func (pm *IPasswordManager) SaveEnterData(login string, data models.EnterData) string {
	data.Key = MakeKey()
	pm.st.SaveEnterData(login, data)
	return data.Key
}

// GetEnterData - получение логина и пароля
func (pm *IPasswordManager) GetEnterData(login, key string) (models.EnterData, error) {
	data, err := pm.st.GetEnterData(login, key)
	return data, err
}

// SaveTextData - сохранение текстовой информации
func (pm *IPasswordManager) SaveTextData(login string, data models.TextData) string {
	data.Key = MakeKey()
	pm.st.SaveTextData(login, data)
	return data.Key
}

// GetTextData - получение текстовой информации
func (pm *IPasswordManager) GetTextData(login, key string) (models.TextData, error) {
	data, err := pm.st.GetTextData(login, key)
	return data, err
}

// SaveBankCardData - сохранение информации о карте
func (pm *IPasswordManager) SaveBankCardData(login string, data models.BankCardData) string {
	data.Key = MakeKey()
	pm.st.SaveBankCardData(login, data)
	return data.Key
}

// GetBankCardData - получение информации о карте
func (pm *IPasswordManager) GetBankCardData(login, key string) (models.BankCardData, error) {
	data, err := pm.st.GetBankCardData(login, key)
	return data, err
}

// GetAllData - получение всей информации
func (pm *IPasswordManager) GetAllData(login string) (models.AllData) {
	data := pm.st.GetAllData(login)
	return data
}

// CheckLogin - проверка наличия логина
func (pm *IPasswordManager) CheckLogin(login string) bool {
	return pm.st.CheckLogin(login)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// MakeKey - создание уникального ключа
func MakeKey() string {
    b := make([]rune, 5)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}