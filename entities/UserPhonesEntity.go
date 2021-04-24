package entities

type UserPhones struct {
	ID        			uint `gorm:"primarykey"`
	UserId			uint
	PhoneNumber			string
	PhoneType			string
}
