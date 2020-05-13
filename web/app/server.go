package app

import (
	"fmt"

	"github.com/PietroCarrara/vtrine/pkg/torrentprovider"
	"github.com/gin-gonic/gin"
)

var provider torrentprovider.TorrentProvider

// Serve starts the server, listening for requests on a port,
// using a torrent provider to fetch torrent data
func Serve(port int, torr torrentprovider.TorrentProvider) {
	loadTemplates()

	provider = torr

	router := gin.Default()

	router.GET("/", index)
	router.Static("/static", "web/static")

	router.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
