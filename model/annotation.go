package model

import "time"

type Annotation struct {
	ID        uint64
	TaskID    uint64
	AuthorId  string `gorm:"not null"`
	Author    User   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Json      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    uint64
}
