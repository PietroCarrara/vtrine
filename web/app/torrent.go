package app

import (
	"encoding/json"
	"net/http"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/gin-gonic/gin"
)

func manage(c *gin.Context) {
	m, _ := client.ListMovies()
	s, _ := client.ListShows()
	media := join(m, s)

	freeSpace, _ := client.GetFreeSpace()
	usedSpace, _ := client.GetUsedSpace()

	err := templates.ExecuteTemplate(c.Writer, "manage.go.html", map[string]interface{}{
		"torrents":  media,
		"freeSpace": float32(freeSpace),
		"usedSpace": float32(usedSpace),
	})

	if err != nil {
		c.Writer.Write([]byte(err.Error()))
	}
}

func download(c *gin.Context) {
	var data torrent.ProviderData
	json.Unmarshal([]byte(c.PostForm("data")), &data)

	switch data.Type {
	case torrent.MediaTypeMovie:
		client.DownloadMovie(data)
	case torrent.MediaTypeTVShow:
		client.DownloadShow(data)
	}

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func remove(c *gin.Context) {
	id := c.PostForm("id")
	client.RemoveTorrent(id)

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func external(c *gin.Context) {
	templates.ExecuteTemplate(c.Writer, "external.go.html", nil)
}

func pause(c *gin.Context) {
	id := c.PostForm("id")
	client.PauseTorrent(id)

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func resume(c *gin.Context) {
	id := c.PostForm("id")
	client.ResumeTorrent(id)

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func join(v ...[]torrent.ClientData) []torrent.ClientData {
	res := make([]torrent.ClientData, 0)

	for _, arr := range v {
		res = append(res, arr...)
	}

	return res
}
