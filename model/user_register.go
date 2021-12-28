package model

type User_register struct {
	U_id int  		    `gorm:"primaryKey;autoIncrement"`
	UserName string     `gorm:"not null"`
	PassWord string     `gorm:"not null"`
	Telephone string    `gorm:"not null"`
	Email string
}
