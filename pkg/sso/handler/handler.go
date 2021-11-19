package handler

import (
	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Login(c *gin.Context)
	Callback(c *gin.Context)
}
