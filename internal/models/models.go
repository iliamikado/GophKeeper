// Пакет с моделями
package models

// Data - инфа с метаинформацией
type Data struct {
	Key      string `json:"-"`
	Metadata string `json:"metadata"`
}

// EnterData - инфа о логине и пароле
type EnterData struct {
	Data
	Login    string `json:"login"`
	Password string `json:"password"`
}

// TextData - инфа с текстом
type TextData struct {
	Data
	Text string `json:"text"`
}

// PaymentCard - инфа о банковской карте
type PaymentCard struct {
	Data
	Number       string `json:"number"`
	YearAndMonth string `json:"year_and_month"`
	CVV          string `json:"cvv"`
}

// AllData - все данные пользователя
type AllData struct {
	EnterData   []EnterData `json:"enter_data"`
	TextData    []TextData `json:"text_data"`
	PaymentCard []PaymentCard `json:"payment_card"`
}
