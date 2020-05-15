package app

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	templates.ExecuteTemplate(c.Writer, "index.go.html", nil)
}
