module github.com/PietroCarrara/vtrine

go 1.14

replace github.com/qopher/go-torrentapi => github.com/PietroCarrara/go-torrentapi v0.1.3

replace github.com/PietroCarrara/flood => ../flood

replace github.com/PietroCarrara/rencode => ../rencode

require (
	github.com/PietroCarrara/flood v0.0.2
	github.com/dustin/go-humanize v1.0.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/hekmon/transmissionrpc v1.1.0
	github.com/joho/godotenv v1.3.0
	github.com/qopher/go-torrentapi v0.0.0-20190920055042-036fa2f8cb12
)
