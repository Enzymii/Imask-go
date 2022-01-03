package model

type Media struct {
	Name    string `gorm:"primary_key"`
	Type    string `gorm:"size:20"`
	OwnerID string `gorm:"size:50"`
	Owner   User   `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
