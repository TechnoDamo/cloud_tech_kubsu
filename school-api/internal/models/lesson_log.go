package models

type LessonLog struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    SubjectID uint   `json:"subject_id" gorm:"not null"`
    Date      string `json:"date" gorm:"type:date;not null"`
    Number    int    `json:"number" gorm:"not null;check:number >= 1 AND number <= 8"`
    ClassID   uint   `json:"class_id" gorm:"not null"`
    TeacherID uint   `json:"teacher_id" gorm:"not null"`
}


