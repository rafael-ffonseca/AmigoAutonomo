package entities

type UserAddress struct {
	ID        	uint `gorm:"primarykey"`
	UserId	uint
	Street		string
	Number		string
	Complement	string
	Region		string
	City		string
	State		string
	Country		string
}
