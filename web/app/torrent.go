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
	a, _ := client.ListAnimes()

	media := join(m, s, a)

	templates.ExecuteTemplate(c.Writer, "manage.go.html", map[string]interface{}{
		"torrents": media,
	})
}

func download(c *gin.Context) {
	var data torrent.ProviderData
	json.Unmarshal([]byte(c.PostForm("data")), &data)

	switch data.Type {
	case torrent.MediaTypeMovie:
		client.DownloadMovie(data)
	case torrent.MediaTypeTVShow:
		client.DownloadShow(data)
	case torrent.MediaTypeAnime:
		client.DownloadAnime(data)
	}

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func remove(c *gin.Context) {
	id := c.PostForm("id")
	client.RemoveTorrent(id)

	c.Redirect(http.StatusFound, "/torrent/manage")
}

func join(v ...[]torrent.ClientData) []torrent.ClientData {
	res := make([]torrent.ClientData, 0)

	for _, arr := range v {
		res = append(res, arr...)
	}

	return res
}
