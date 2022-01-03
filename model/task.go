package model

import "time"

type Task struct {
	ID        uint64
	Name      string `gorm:"not null"`
	AuthorId  string `gorm:"not null"`
	Author    User   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}
