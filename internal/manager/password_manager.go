// Пакет с основной логикой
package manager

import (
	"PasswordManager/internal/models"
	"PasswordManager/internal/storage"
	"math/rand"
)

// PasswordManager - струкура основного приложения
type PasswordManager struct {
	st storage.Storage
}

// PasswordManager - основнай экзэмпляр реализующий логику
var AppManager *PasswordManager

// Initialize - инициализация PasswordManager
func Initialize() {
	AppManager = &PasswordManager{
		st: storage.NewSimpleStorage(),
	}
}

// Register - регистрация пользователя
func (pm *PasswordManager) Register(login, password string) (string, error) {
	err := pm.st.Register(login, password)
	if err != nil {
		return "", err
	}
	return login, nil
}

// Login - вход пользователя
func (pm *PasswordManager) Login(login, password string) (string, error) {
	err := pm.st.Login(login, password)
	if err != nil {
		return "", err
	}
	return login, nil
}

// SaveEnterData - сохранение логина и пароля
func (pm *PasswordManager) SaveEnterData(login string, data models.EnterData) string {
	data.Key = MakeKey()
	pm.st.SaveEnterData(login, data)
	return data.Key
}

// GetEnterData - получение логина и пароля
func (pm *PasswordManager) GetEnterData(login, key string) (models.EnterData, error) {
	data, err := pm.st.GetEnterData(login, key)
	return data, err
}

// SaveTextData - сохранение текстовой информации
func (pm *PasswordManager) SaveTextData(login string, data models.TextData) string {
	data.Key = MakeKey()
	pm.st.SaveTextData(login, data)
	return data.Key
}

// GetTextData - получение текстовой информации
func (pm *PasswordManager) GetTextData(login, key string) (models.TextData, error) {
	data, err := pm.st.GetTextData(login, key)
	return data, err
}

// SavePaymentCard - сохранение информации о карте
func (pm *PasswordManager) SavePaymentCard(login string, data models.PaymentCard) string {
	data.Key = MakeKey()
	pm.st.SavePaymentCard(login, data)
	return data.Key
}

// GetPaymentCard - получение информации о карте
func (pm *PasswordManager) GetPaymentCard(login, key string) (models.PaymentCard, error) {
	data, err := pm.st.GetPaymentCard(login, key)
	return data, err
}

// GetAllData - получение всей информации
func (pm *PasswordManager) GetAllData(login string) models.AllData {
	data := pm.st.GetAllData(login)
	return data
}

// CheckLogin - проверка наличия логина
func (pm *PasswordManager) CheckLogin(login string) bool {
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
