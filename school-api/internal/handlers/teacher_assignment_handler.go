package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "school-api/internal/models"
)

type TeacherAssignmentHandler struct{ DB *gorm.DB }

func (h TeacherAssignmentHandler) Register(r *gin.RouterGroup) {
    r.GET("/teacher-assignments", h.List)
    r.POST("/teacher-assignments", h.Create)
    r.GET("/teacher-assignments/:id", h.Get)
    r.PUT("/teacher-assignments/:id", h.Update)
    r.DELETE("/teacher-assignments/:id", h.Delete)
}

func (h TeacherAssignmentHandler) List(c *gin.Context) {
    var items []models.TeacherAssignment
    if err := h.DB.Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": items})
}

func (h TeacherAssignmentHandler) Create(c *gin.Context) {
    var input models.TeacherAssignment
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

func (h TeacherAssignmentHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.TeacherAssignment
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

func (h TeacherAssignmentHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var item models.TeacherAssignment
    if err := h.DB.First(&item, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "NotFound", "message": "Resource not found"})
        return
    }
    var input models.TeacherAssignment
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "BadRequest", "message": err.Error()})
        return
    }
    item.TeacherID = input.TeacherID
    item.SubjectID = input.SubjectID
    if err := h.DB.Save(&item).Error; err != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "UnprocessableEntity", "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, item)
}

func (h TeacherAssignmentHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.DB.Delete(&models.TeacherAssignment{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "InternalServerError", "message": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


