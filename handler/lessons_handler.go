package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLesson(c *gin.Context) {
	var body models.Lesson
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lesson, err := h.storage.CreateLesson(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

func (h *Handler) GetLessons(c *gin.Context) {
	lessons, err := h.storage.GetLessons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func (h *Handler) GetLesson(c *gin.Context) {
	lessonIdStr := c.Param("lesson_id")
	lessonId, err := strconv.Atoi(lessonIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lesson, err := h.storage.GetLesson(lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lesson)

}


func (h *Handler) PutLesson(c *gin.Context) {
	lessonIdStr := c.Param("lesson_id")
	lessonId, err := strconv.Atoi(lessonIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var body models.Lesson
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lesson, err := h.storage.PutLesson(lessonId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lesson)
}



func (h *Handler) CompleteLesson(c *gin.Context) {
	lessonIdStr := c.Param("lesson_id")
	lessonId, err := strconv.Atoi(lessonIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	lesson, err := h.storage.CompleteLesson(lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	lessonIdStr := c.Param("lesson_id")
	lessonId, err := strconv.Atoi(lessonIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.storage.DeleteLesson(lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "The lesson was deleted successfully!!!",
	})
}
