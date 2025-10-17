package models

type Class struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Grade  int    `json:"grade" gorm:"not null;check:grade >= 1 AND grade <= 12"`
	Letter string `json:"letter" gorm:"type:char(1);not null;check:letter ~ '^[A-Z]'"`
}
