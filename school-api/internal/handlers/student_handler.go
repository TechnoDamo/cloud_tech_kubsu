package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type StudentHandler struct{ DB *gorm.DB }

func (h StudentHandler) Register(r *gin.RouterGroup) {
    r.GET("/students", h.List)
    r.POST("/students", h.Create)
    r.GET("/students/:id", h.Get)
    r.PUT("/students/:id", h.Update)
    r.DELETE("/students/:id", h.Delete)
}

func (h StudentHandler) List(c *gin.Context) {
    var items []models.Student
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h StudentHandler) Create(c *gin.Context) {
    var input models.Student
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

func (h StudentHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.Student
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

func (h StudentHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.Student
    if err := h.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.Student
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.ClassID = input.ClassID
    item.FirstName = input.FirstName
    item.LastName = input.LastName
    item.Patronymic = input.Patronymic
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h StudentHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.Student{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


