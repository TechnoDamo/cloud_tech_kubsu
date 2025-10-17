package models

type Student struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    ClassID    uint   `json:"class_id" gorm:"not null"`
    FirstName  string `json:"first_name" gorm:"type:varchar(50);not null"`
    LastName   string `json:"last_name" gorm:"type:varchar(50);not null"`
    Patronymic string `json:"patronymic" gorm:"type:varchar(50)"`
}


