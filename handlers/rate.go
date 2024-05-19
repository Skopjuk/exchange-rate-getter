package handlers

import (
	"awesomeProject7/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getRate(c *gin.Context) {
	rate, err := services.GetExchangeRate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status value"})
	}
	c.JSON(http.StatusOK, rate)
}
