package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func download(c *gin.Context) {
	mag := c.PostForm("magnet")
	category := c.PostForm("category")

	switch category {
	case "movie":
		client.DownloadMovie(mag)
	case "tv":
		client.DownloadShow(mag)
	case "anime":
		client.DownloadAnime(mag)
	}

	c.Redirect(http.StatusFound, "/")
}
