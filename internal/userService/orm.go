package userService

import "gorm.io/gorm"

// Message - структура, представляющая сообщение в БД.
type User struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      int    `gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
