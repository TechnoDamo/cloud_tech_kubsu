package models

type StudentLesson struct {
    ID               uint   `json:"id" gorm:"primaryKey"`
    StudentID        uint   `json:"student_id" gorm:"not null"`
    LessonID         uint   `json:"lesson_id" gorm:"not null"`
    Grade            *int   `json:"grade"`
    AttendanceStatus string `json:"attendance_status" gorm:"type:char(1);not null"`
}


