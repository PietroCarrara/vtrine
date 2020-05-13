package rarbg

import (
	"strings"

	"github.com/PietroCarrara/vtrine/pkg/torrentprovider"
	"github.com/qopher/go-torrentapi"
)

func data(res torrentapi.TorrentResults) []torrentprovider.TorrentData {
	torrents := make([]torrentprovider.TorrentData, len(res))

	for i, torr := range res {
		torrents[i] = torrentprovider.TorrentData{
			Title:  torr.Title,
			Type:   mediaType(torr.Category),
			Magnet: torr.Download,
			Size:   torr.Size,
			IMDB:   torr.EpisodeInfo.ImDB,
			TMDB:   torr.EpisodeInfo.TheMovieDb,
		}

		// Clean bad data
		if torrents[i].IMDB == "0" {
			torrents[i].IMDB = ""
		}
		if torrents[i].TMDB == "0" {
			torrents[i].TMDB = ""
		}
	}

	return torrents
}

func mediaType(category string) torrentprovider.MediaType {
	low := strings.ToLower(category)

	if strings.HasPrefix(low, "movies") {
		return torrentprovider.MediaTypeMovie
	} else if strings.HasPrefix(low, "tv") {
		return torrentprovider.MediaTypeTVShow
	} else {
		return torrentprovider.MediaTypeUnknown
	}
}

func empty() []torrentprovider.TorrentData {
	return []torrentprovider.TorrentData{}
}
