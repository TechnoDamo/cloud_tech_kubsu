package models

type TeacherAssignment struct {
    ID        uint `json:"id" gorm:"primaryKey"`
    TeacherID uint `json:"teacher_id" gorm:"not null"`
    SubjectID uint `json:"subject_id" gorm:"not null"`
}


