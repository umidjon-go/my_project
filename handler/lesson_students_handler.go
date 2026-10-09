package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLessonStudent(c *gin.Context) {
	var body models.Lesson_Student
	if err := c.ShouldBindJSON(&body); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lessonStudent, err := h.storage.CreateLessonStudent(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, lessonStudent)
}
func (h *Handler) GetLessonStudent(c *gin.Context) {
	lessStuIDStr := c.Param("lesson_student_id")
	lessonStuID, err := strconv.Atoi(lessStuIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lessonStudent, err := h.storage.GetLessonStudent(lessonStuID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lessonStudent)
}