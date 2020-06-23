package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/PietroCarrara/vtrine/pkg/torrent/deluge"
	"github.com/PietroCarrara/vtrine/pkg/torrent/rarbg"
	"github.com/PietroCarrara/vtrine/pkg/torrent/transmission"
	"github.com/PietroCarrara/vtrine/web/app"
	"github.com/joho/godotenv"
)

func fail(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	godotenv.Load()

	var client torrent.TorrentClient
	var err error

	switch os.Getenv("CLIENT") {
	case "deluge":
		client, err = deluge.New()
	case "transmission":
		client, err = transmission.New()
	default:
		err = fmt.Errorf("unknown client %s", os.Getenv("CLIENT"))
	}

	fail(err)

	rar, err := rarbg.New("vtrine")
	fail(err)

	app.Serve(54321, rar, client)
}
