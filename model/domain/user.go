package domain

import "time"

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	Role  	  string `gorm:"column:role"`
	CreatedAt time.Time 
}

func (User) TableName() string {
	return "users"
}