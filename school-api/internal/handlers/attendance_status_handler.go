package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type AttendanceStatusHandler struct{ DB *gorm.DB }

func (h AttendanceStatusHandler) Register(r *gin.RouterGroup) {
    r.GET("/attendance-statuses", h.List)
    r.POST("/attendance-statuses", h.Create)
    r.GET("/attendance-statuses/:code", h.Get)
    r.PUT("/attendance-statuses/:code", h.Update)
    r.DELETE("/attendance-statuses/:code", h.Delete)
}

func (h AttendanceStatusHandler) List(c *gin.Context) {
    var items []models.AttendanceStatus
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h AttendanceStatusHandler) Create(c *gin.Context) {
    var input models.AttendanceStatus
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    if err := h.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, input)
}

func (h AttendanceStatusHandler) Get(c *gin.Context) {
    code := c.Param("code")
    var item models.AttendanceStatus
    if err := h.DB.First(&item, "code = ?", code).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h AttendanceStatusHandler) Update(c *gin.Context) {
    code := c.Param("code")
    var item models.AttendanceStatus
    if err := h.DB.First(&item, "code = ?", code).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.AttendanceStatus
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.Description = input.Description
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h AttendanceStatusHandler) Delete(c *gin.Context) {
    code := c.Param("code")
    if err := h.DB.Delete(&models.AttendanceStatus{}, "code = ?", code).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


