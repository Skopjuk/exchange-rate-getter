package handlers

import (
	"awesomeProject7/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *services.Service
	logger   logrus.Logger
}

func NewHandler(services *services.Service, logger logrus.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.GET("/rate", h.getRate)
	router.POST("/subscribe", h.createSubscription)

	return router
}
