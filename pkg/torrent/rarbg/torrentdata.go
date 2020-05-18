package rarbg

import (
	"strings"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/qopher/go-torrentapi"
)

func data(res torrentapi.TorrentResults) []torrent.TorrentData {
	torrents := make([]torrent.TorrentData, len(res))

	for i, torr := range res {
		torrents[i] = torrent.TorrentData{
			Title:        torr.Title,
			Type:         mediaType(torr.Category),
			Magnet:       torr.Download,
			Size:         torr.Size,
			IMDB:         torr.EpisodeInfo.ImDB,
			ProviderName: name,
		}

		// Clean bad data
		if torrents[i].IMDB == "0" {
			torrents[i].IMDB = ""
		}
	}

	return torrents
}

func mediaType(category string) torrent.MediaType {
	low := strings.ToLower(category)

	if strings.Contains(low, "tv") {
		return torrent.MediaTypeTVShow
	} else if strings.Contains(low, "movies") {
		return torrent.MediaTypeMovie
	} else {
		return torrent.MediaTypeUnknown
	}
}

func empty() []torrent.TorrentData {
	return []torrent.TorrentData{}
}
