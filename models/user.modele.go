package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
