package controller

import (
	"cs-lab-6/web/index"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleIndex(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	c.Writer.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(c.Writer, "%s", index.IndexPage)
}
