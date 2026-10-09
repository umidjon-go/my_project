package handler

import (
	"edu-crm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	var body models.Payment
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	payment, err := h.storage.CreatePayment(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, payment)
}
func (h *Handler) GetPayment(c *gin.Context) {
	paymentIdStr := c.Param("payment_id")
	paymentId, err := strconv.Atoi(paymentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	payment, err := h.storage.GetPayment(paymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, payment)
}
func (h *Handler) PutPayment(c *gin.Context) {
	paymentIdStr := c.Param("payment_id")
	paymentId, err := strconv.Atoi(paymentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	payment, err := h.storage.PutPayment(paymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *Handler) DeletePayment(c *gin.Context) {
	paymentIdStr := c.Param("payment_id")
	paymentId, err := strconv.Atoi(paymentIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.storage.DeletePayment(paymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "The payment was deleted successfully!!!",
	})
}