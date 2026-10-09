package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStudent(c *gin.Context) {
	var body models.Student
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	student, err := h.storage.CreateStudent(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, student)
}
func (h *Handler) GetStudent(c *gin.Context) {
	studentIdStr := c.Param("student_id")
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	student, err := h.storage.GetStudent(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
func (h *Handler) PutStudent(c *gin.Context) {
	studentIdStr := c.Param("student_id")
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	student, err := h.storage.PutStudent(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
func (h *Handler) DeleteStudent(c *gin.Context) {
	studentIdStr := c.Param("student_id")
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.storage.DeleteStudent(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "the student was deleted successfully!!!",
	})
}	