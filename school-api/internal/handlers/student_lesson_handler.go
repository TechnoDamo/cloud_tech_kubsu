package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type StudentLessonHandler struct{ DB *gorm.DB }

func (h StudentLessonHandler) Register(r *gin.RouterGroup) {
    r.GET("/student-lessons", h.List)
    r.POST("/student-lessons", h.Create)
    r.GET("/student-lessons/:id", h.Get)
    r.PUT("/student-lessons/:id", h.Update)
    r.DELETE("/student-lessons/:id", h.Delete)
}

func (h StudentLessonHandler) List(c *gin.Context) {
    var items []models.StudentLesson
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h StudentLessonHandler) Create(c *gin.Context) {
    var input models.StudentLesson
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    input.ID = 0
    if err := h.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, input)
}

func (h StudentLessonHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.StudentLesson
    if err := h.DB.First(&item, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h StudentLessonHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.StudentLesson
    if err := h.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.StudentLesson
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.StudentID = input.StudentID
    item.LessonID = input.LessonID
    item.Grade = input.Grade
    item.AttendanceStatus = input.AttendanceStatus
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h StudentLessonHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.StudentLesson{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


