package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type TeacherHandler struct{ DB *gorm.DB }

func (h TeacherHandler) Register(r *gin.RouterGroup) {
    r.GET("/teachers", h.List)
    r.POST("/teachers", h.Create)
    r.GET("/teachers/:id", h.Get)
    r.PUT("/teachers/:id", h.Update)
    r.DELETE("/teachers/:id", h.Delete)
}

func (h TeacherHandler) List(c *gin.Context) {
    var items []models.Teacher
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h TeacherHandler) Create(c *gin.Context) {
    var input models.Teacher
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

func (h TeacherHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.Teacher
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

func (h TeacherHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.Teacher
    if err := h.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.Teacher
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.FirstName = input.FirstName
    item.LastName = input.LastName
    item.Patronymic = input.Patronymic
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h TeacherHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.Teacher{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


