package models

type AttendanceStatus struct {
    Code        string `json:"code" gorm:"type:char(1);primaryKey"`
    Description string `json:"description" gorm:"type:varchar(50);not null"`
}


