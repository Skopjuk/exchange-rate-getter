package handlers

import (
	"awesomeProject7/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSubscription(c *gin.Context) {
	var input models.Subscription

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error()})
		return
	}

	err := h.services.Subscription.GetSubscriptions(input.Email)
	if err == nil {
		c.JSON(http.StatusConflict, map[string]interface{}{
			"error": "subscription already exists"})
		return
	}

	err = h.services.Subscription.CreateSubscription(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
