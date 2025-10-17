package models

type Subject struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    SubjectName string `json:"subject_name" gorm:"type:varchar(100);not null"`
}


