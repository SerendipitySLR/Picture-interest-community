package model

import "time"

type User_detail struct {
	U_id int           			`gorm:"primaryKey;autoIncrement"`
	Nickname string				`gorm:"not null"`
	Sex bool					`gorm:"not null"`
	Register_data time.Time		`gorm:"not null"`
	Update_date time.Time		`gorm:"not null"`
	Birthday time.Time
	Location string
	Signature string
	Profile_url string
}
