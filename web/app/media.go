package app

import "github.com/gin-gonic/gin"

func media(c *gin.Context) {
	imdb := c.Param("imdb")
	torrents, _ := provider.SearchIMDB(imdb)

	templates.ExecuteTemplate(c.Writer, "media.go.html", map[string]interface{}{
		"imdb":     imdb,
		"torrents": torrents,
	})
}
