package auth

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"unique;not null;size:64"`
	PasswordHash string    `json:"-" gorm:"not null;size:128"`
	Name         string    `json:"name" gorm:"size:128"`
	Email        string    `json:"email" gorm:"size:128"`
	Role         string    `json:"role" gorm:"size:32;default:it_support"`
	CreatedAt    time.Time `json:"created_at"`
}
