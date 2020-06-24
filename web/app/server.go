package app

import (
	"fmt"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/gin-gonic/gin"
)

var provider torrent.TorrentProvider
var client torrent.TorrentClient

// Serve starts the server, listening for requests on a port,
// using a torrent provider to fetch torrent data
func Serve(port int, torr torrent.TorrentProvider, cli torrent.TorrentClient) {
	loadTemplates()

	provider = torr
	client = cli

	router := gin.Default()

	router.GET("/", index)
	router.GET("/media/:imdb", media)
	router.GET("/search/name", searchName)
	router.GET("/torrent/manage", manage)
	router.GET("/torrent/external", external)
	router.POST("/torrent/download", download)
	router.POST("/torrent/delete", remove)
	router.Static("/static", "web/static")

	router.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
