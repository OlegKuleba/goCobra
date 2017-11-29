package models

// Структура контакта
type Contact struct {
	Phone   string
	Name    string
	Address string
}

// Конструктор для контакта
func NewContact(phone string, name string, address string) Contact {
	return Contact{
		Phone:   phone,
		Name:    name,
		Address: address,
	}
}