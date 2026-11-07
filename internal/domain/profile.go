package domain

import "time"

// Profile là hồ sơ của một admin backoffice.
type Profile struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	Email      string
	Department string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Profile) TableName() string { return "profiles" }
