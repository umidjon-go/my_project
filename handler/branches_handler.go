package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBranch(c *gin.Context) {
	var body models.Branch
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	branch, err := h.storage.CreateBranch(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, branch)
}
func (h *Handler) GetBranch(c *gin.Context) {
	branchIdstr := c.Param("branch_id")
	branchId, err := strconv.Atoi(branchIdstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	branch, err := h.storage.GetBranch(branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, branch)
}
func (h *Handler) GetBranches(c *gin.Context) {
	branches, err := h.storage.GetBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, branches)
}
func (h *Handler) PutBranch(c *gin.Context) {
	branchIdstr := c.Param("branch_id")
	branchId, err := strconv.Atoi(branchIdstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var body models.Branch
	if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	branch, err := h.storage.PutBranch(branchId, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, branch)
}
func (h *Handler) DeleteBranch(c *gin.Context) {
	branchIdstr := c.Param("branch_id")
	branchId, err := strconv.Atoi(branchIdstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.storage.DeleteBranch(branchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Delete": "The branch was deleted successfully!!!",
	})
}