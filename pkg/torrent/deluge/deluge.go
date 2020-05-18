package deluge

import (
	"os"

	"github.com/PietroCarrara/flood"
)

type Deluge struct {
	*flood.Flood
}

func New() (*Deluge, error) {
	f, err := flood.New(os.Getenv("DELUGE_ADDR"))
	if err != nil {
		return nil, err
	}

	d := Deluge{f}
	_, err = d.Login(os.Getenv("DELUGE_USR"), os.Getenv("DELUGE_PASS"))

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (d *Deluge) DownloadMovie(magnet string) error {
	return d.downloadLabel(magnet, "movies")
}

func (d *Deluge) DownloadShow(magnet string) error {
	return d.downloadLabel(magnet, "shows")
}

func (d *Deluge) DownloadAnime(magnet string) error {
	return d.downloadLabel(magnet, "animes")
}

func (d *Deluge) downloadLabel(magnet, label string) error {
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

	id, err := d.Core.AddTorrentMagnet(magnet)
	if err != nil {
		return err
	}

	err = d.Label.SetTorrent(id, label)
	return err
}

func contains(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}

	return false
}
