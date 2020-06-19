package app

import (
	"strings"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
)

const gigabyte = 1024 * 1024 * 1025

// bestTorrent selects the best torrent of a media.
// When deploying, users might wanna modify this function
func bestTorrent(t []torrent.ProviderData) *torrent.ProviderData {
	for _, t := range t {
		name := strings.ToLower(t.Title)

		if strings.Contains(name, "264") && strings.Contains(name, "1080p") && t.Size < 6*gigabyte && t.Type == torrent.MediaTypeMovie {
			return &t
		}
	}

	return nil
}
