package entities

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;autoIngrement:true" json:"id"`
	Name      string    `gorm:"type:varchar(50); not null;" json:"name"`
	CreatedAt time.Time `json:"-"`
}

func NewUser(id uint, name string) *User {
	user := &User{
		ID:   id,
		Name: name,
	}
	return user
}
