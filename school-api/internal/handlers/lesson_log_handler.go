package handlers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type LessonLogHandler struct{ DB *gorm.DB }

func (h LessonLogHandler) Register(r *gin.RouterGroup) {
    r.GET("/lesson-logs", h.List)
    r.POST("/lesson-logs", h.Create)
    r.GET("/lesson-logs/:id", h.Get)
    r.PUT("/lesson-logs/:id", h.Update)
    r.DELETE("/lesson-logs/:id", h.Delete)
}

func (h LessonLogHandler) List(c *gin.Context) {
    var items []models.LessonLog
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h LessonLogHandler) Create(c *gin.Context) {
    var input models.LessonLog
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

func (h LessonLogHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.LessonLog
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

func (h LessonLogHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.LessonLog
    if err := h.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.LessonLog
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.SubjectID = input.SubjectID
    item.Date = input.Date
    item.Number = input.Number
    item.ClassID = input.ClassID
    item.TeacherID = input.TeacherID
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h LessonLogHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.LessonLog{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


