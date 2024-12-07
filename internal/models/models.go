// Пакет с моделями
package models

// Data - инфа с метаинформацией
type Data struct {
	Key      string
	Metadata string
}

// EnterData - инфа о логине и пароле
type EnterData struct {
	Data
	Login    string
	Password string
}

// TextData - инфа с текстом
type TextData struct {
	Data
	Text string
}

// PaymentCard - инфа о банковской карте
type PaymentCard struct {
	Data
	Number       string
	YearAndMonth string
	CVV          string
}

// AllData - все данные пользователя
type AllData struct {
	EnterData   []EnterData
	TextData    []TextData
	PaymentCard []PaymentCard
}
