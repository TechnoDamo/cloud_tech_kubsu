package models

type LessonSchedule struct {
    ID        uint `json:"id" gorm:"primaryKey"`
    SubjectID uint `json:"subject_id" gorm:"not null"`
    Weekday   int  `json:"weekday" gorm:"not null;check:weekday >= 1 AND weekday <= 7"`
    Number    int  `json:"number" gorm:"not null;check:number >= 1 AND number <= 8"`
    ClassID   uint `json:"class_id" gorm:"not null"`
    TeacherID uint `json:"teacher_id" gorm:"not null"`
}


