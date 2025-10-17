package router

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/handlers"
)

func Setup(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    
    // Add CORS middleware
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Header("Access-Control-Allow-Credentials", "true")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })
    
    api := r.Group("/api/v1")

    handlers.ClassHandler{DB: db}.Register(api)
    handlers.StudentHandler{DB: db}.Register(api)
    handlers.TeacherHandler{DB: db}.Register(api)
    handlers.SubjectHandler{DB: db}.Register(api)
    handlers.TeacherAssignmentHandler{DB: db}.Register(api)
    handlers.LessonScheduleHandler{DB: db}.Register(api)
    handlers.LessonLogHandler{DB: db}.Register(api)
    handlers.StudentLessonHandler{DB: db}.Register(api)
    handlers.AttendanceStatusHandler{DB: db}.Register(api)
    return r
}

