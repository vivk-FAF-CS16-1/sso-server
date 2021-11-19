package controller

import (
	"cs-lab-6/pkg/sso/handler"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(r *gin.RouterGroup, handler handler.IHandler) {
	r.GET("/login", handler.Login)
	r.GET("/callback", handler.Callback)
}
