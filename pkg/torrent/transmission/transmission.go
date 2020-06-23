package transmission

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/hekmon/transmissionrpc"
)

const (
	movieLocation = "movies"
	showLocation  = "shows"
)

type Transmission struct {
	c                *transmissionrpc.Client
	downloadLocation string
}

func New() (*Transmission, error) {
	addr := os.Getenv("TRANSMISSION_ADDR")
	if addr == "" {
		return nil, errors.New("TRANSMISSION_ADDR environment variable not set")
	}

	location := os.Getenv("TRANSMISSION_DOWNLOAD")
	if location == "" {
		return nil, errors.New("TRANSMISSION_DOWNLOAD environment variable not set")
	}
	if !strings.HasSuffix(location, "/") {
		location += "/"
	}

	user := os.Getenv("TRANSMISSION_USER")
	pass := os.Getenv("TRANSMISSION_PASSWORD")

	c, err := transmissionrpc.New(addr, user, pass, nil)
	if err != nil {
		return nil, err
	}

	return &Transmission{
		c:                c,
		downloadLocation: location,
	}, nil
}

// DownloadMovie downloads a magnet link to the movies folder
func (t *Transmission) DownloadMovie(data torrent.ProviderData) error {
	return t.downloadToLocation(data, movieLocation)
}

// DownloadShow downloads a magnet link to the tv folder
func (t *Transmission) DownloadShow(data torrent.ProviderData) error {
	return t.downloadToLocation(data, showLocation)
}

// ListMovies lists all the movies registered within the movies folder
func (t *Transmission) ListMovies() ([]torrent.ClientData, error) {
	return t.listInLocation(movieLocation)
}

// ListShows lists all the movies registered within the tv folder
func (t *Transmission) ListShows() ([]torrent.ClientData, error) {
	return t.listInLocation(showLocation)
}

// GetFreeSpace returns the free space in the disk
func (t *Transmission) GetFreeSpace() (uint64, error) {
	bytes, err := t.c.FreeSpace(t.downloadLocation)
	if err != nil {
		return 0, err
	}

	return uint64(bytes.Byte()), nil
}

func (t *Transmission) GetUsedSpace() (uint64, error) {
	torrents, err := t.c.TorrentGet([]string{
		"totalSize",
	}, nil)

	if err != nil {
		return 0, err
	}

	var total uint64
	for _, torr := range torrents {
		total += uint64(torr.TotalSize.Byte())
	}

	return total, nil
}

// RemoveTorrent removes a torrent alongside it's data
func (t *Transmission) RemoveTorrent(id string) error {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	err = t.c.TorrentRemove(&transmissionrpc.TorrentRemovePayload{
		IDs:             []int64{idInt},
		DeleteLocalData: true,
	})

	return err
}

// Download a torrent to a location inside t.downloadLocation
func (t *Transmission) downloadToLocation(data torrent.ProviderData, location string) error {
	path := fmt.Sprintf("%s%s/%s.%s/", t.downloadLocation, location, data.Title, data.IMDB)
	_, err := t.c.TorrentAdd(&transmissionrpc.TorrentAddPayload{
		Filename:    &data.Magnet,
		DownloadDir: &path,
	})
	return err
}

func (t *Transmission) listInLocation(location string) ([]torrent.ClientData, error) {
	res := make([]torrent.ClientData, 0)

	torrents, err := t.c.TorrentGet([]string{
		"id",
		"downloadDir",
		"name",
		"doneDate",
		"percentDone",
		"totalSize",
		"status",
	}, nil)

	if err != nil {
		return nil, err
	}

	for _, torr := range torrents {
		if strings.HasPrefix(*torr.DownloadDir, t.downloadLocation+location) {
			res = append(res, clientDataFromTorrent(torr))
		}
	}

	return res, nil
}

func clientDataFromTorrent(t *transmissionrpc.Torrent) torrent.ClientData {
	id := fmt.Sprint(t.ID)

	// Split the directories, and get the imdb after the last '.'
	parts := strings.FieldsFunc(*t.DownloadDir, func(r rune) bool {
		return r == '/'
	})
	parts = strings.FieldsFunc(parts[len(parts)-1], func(r rune) bool {
		return r == '.'
	})
	imdb := parts[len(parts)-1]

	return torrent.ClientData{
		ID:       id,
		Complete: t.DoneDate.IsZero(),
		IMDB:     imdb,
		Paused:   *t.Status != transmissionrpc.TorrentStatusDownload,
		Progress: float32(*t.PercentDone),
		Size:     uint64(t.TotalSize.Byte()),
		Title:    *t.Name,
	}
}
