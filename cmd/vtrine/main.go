package main

import (
	"log"

	"github.com/PietroCarrara/vtrine/pkg/torrent/deluge"
	"github.com/PietroCarrara/vtrine/pkg/torrent/rarbg"
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

	rar, err := rarbg.New("vtrine")
	fail(err)

	del, err := deluge.New()
	fail(err)

	app.Serve(54321, rar, del)
}
