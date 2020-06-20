package deluge

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/PietroCarrara/flood"
	"github.com/PietroCarrara/vtrine/pkg/torrent"
)

type Deluge struct {
	*flood.Flood
	DownloadLocation string
}

const (
	movieLabel  = "movies"
	showsLabel  = "shows"
	animesLabel = "animes"
)

func New() (*Deluge, error) {
	addr := os.Getenv("DELUGE_ADDR")
	if addr == "" {
		return nil, errors.New("DELUGE_ADDR environment variable not set")
	}

	location := os.Getenv("DELUGE_DOWNLOAD")
	if location == "" {
		return nil, errors.New("DELUGE_DOWNLOAD environment variable not set")
	}
	if !strings.HasSuffix(location, "/") {
		location += "/"
	}

	f, err := flood.New(addr)
	if err != nil {
		return nil, err
	}

	d := Deluge{f, location}
	_, err = d.Login(os.Getenv("DELUGE_USR"), os.Getenv("DELUGE_PASS"))

	if err != nil {
		return nil, err
	}

	return &d, nil
}

// DownloadMovie downloads a magnet link and adds it to the movies label
func (d *Deluge) DownloadMovie(data torrent.ProviderData) error {
	return d.downloadLabel(data, movieLabel)
}

// DownloadShow downloads a magnet link and adds it to the tv label
func (d *Deluge) DownloadShow(data torrent.ProviderData) error {
	return d.downloadLabel(data, showsLabel)
}

// DownloadAnime downloads a magnet link and adds it to the anime label
func (d *Deluge) DownloadAnime(data torrent.ProviderData) error {
	return d.downloadLabel(data, animesLabel)
}

// ListMovies lists all the movies registered within the movies label
func (d *Deluge) ListMovies() ([]torrent.ClientData, error) {
	return d.listLabel(movieLabel)
}

// ListShows lists all the movies registered within the tv label
func (d *Deluge) ListShows() ([]torrent.ClientData, error) {
	return d.listLabel(showsLabel)
}

// ListAnimes lists all the movies registered within the anime label
func (d *Deluge) ListAnimes() ([]torrent.ClientData, error) {
	return d.listLabel(animesLabel)
}

// RemoveTorrent removes a torrent alongside it's data
func (d *Deluge) RemoveTorrent(id string) error {
	_, err := d.Core.RemoveTorrent(id, true)
	return err
}

func (d *Deluge) downloadLabel(data torrent.ProviderData, label string) error {
	labels, err := d.Flood.Label.GetLabels()
	if err != nil {
		return err
	}

	if !contains(labels, label) {
		err = d.Label.Add(label)
		if err != nil {
			return err
		}
	}

	id, err := d.Core.AddTorrentMagnetOptions(data.Magnet, map[string]interface{}{
		flood.MoveOnCompletedField:     true,
		flood.MoveOnCompletedPathField: fmt.Sprintf("%s%s/%s.%s/", d.DownloadLocation, label, data.Title, data.IMDB),
	})
	if err != nil {
		return err
	}

	err = d.Label.SetTorrent(id, label)
	return err
}

func (d *Deluge) listLabel(label string) ([]torrent.ClientData, error) {
	data, err := d.Flood.Core.GetTorrentsStatus(
		map[string]interface{}{
			"label": label,
		},
		flood.TimeAddedField, // Sort By date added
		flood.MoveOnCompletedPathField,
		flood.HashField,
		flood.NameField,
		flood.IsFinishedField,
		flood.PausedField,
		flood.TotalSizeField,
		flood.ProgressField,
	)

	if err != nil {
		return nil, err
	}

	res := clientDataFromStatus(data)

	return res, nil
}

func clientDataFromStatus(statuses []flood.TorrentStatus) []torrent.ClientData {
	res := make([]torrent.ClientData, 0)

	for _, status := range statuses {
		// Split the directories, and get the imdb after the last '.'
		parts := strings.FieldsFunc(status.MoveOnCompletedPath, func(r rune) bool {
			return r == '/'
		})
		parts = strings.FieldsFunc(parts[len(parts)-1], func(r rune) bool {
			return r == '.'
		})
		imdb := parts[len(parts)-1]

		res = append(res, torrent.ClientData{
			ID:       status.Hash,
			Title:    status.Name,
			Complete: status.Finished,
			Size:     status.TotalSize,
			Paused:   status.Paused,
			Progress: status.Progress / 100,
			IMDB:     imdb,
		})
	}

	return res
}

func contains(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}
