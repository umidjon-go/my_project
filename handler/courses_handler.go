package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	var body models.Course
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	course, err := h.storage.CreateCourse(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, course)
}
func (h *Handler) GetCourse(c *gin.Context) {
	courseIdStr := c.Param("course_id")
	courseId, err := strconv.Atoi(courseIdStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	course, err := h.storage.GetCourse(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, course)
}
func (h *Handler) PutCourse(c *gin.Context) {
	courseIdStr := c.Param("course_id")
	courseId, err := strconv.Atoi(courseIdStr)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	course, err := h.storage.PutCourse(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, course)
}
func (h *Handler) DeleteCourse(c *gin.Context) {
	courseIdStr := c.Param("course_id")
	courseId, err := strconv.Atoi(courseIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.storage.DeleteCourse(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
       "delete": "The course was deleted successfully!!!",
	})
}