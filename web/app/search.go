package app

import "github.com/gin-gonic/gin"

func searchName(c *gin.Context) {
	templates.ExecuteTemplate(c.Writer, "search.go.html", map[string]string{
		"query": c.Query("q"),
	})
}
