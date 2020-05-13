package app

import (
	"github.com/PietroCarrara/vtrine/pkg/torrentprovider"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	new, err := provider.New()
	fail(err)

	err = templates.ExecuteTemplate(c.Writer, "index.go.html", map[string]interface{}{
		"new": uniqueTMDB(new),
	})
	fail(err)
}

func uniqueTMDB(torrents []torrentprovider.TorrentData) []torrentprovider.TorrentData {
	present := make(map[string]bool)
	res := make([]torrentprovider.TorrentData, 0)

	for _, torr := range torrents {
		if !present[torr.IMDB] {
			res = append(res, torr)
			present[torr.IMDB] = true
		}
	}

	return res
}
