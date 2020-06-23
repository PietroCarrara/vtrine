package rarbg

import (
	"github.com/PietroCarrara/vtrine/pkg/torrent"
	"github.com/qopher/go-torrentapi"
)

const format = "json_extended"
const name = "rargb"

type Rarbg struct {
	api *torrentapi.API
}

func New(appName string) (*Rarbg, error) {
	api, err := torrentapi.New(appName)

	return &Rarbg{
		api: api,
	}, err
}

func (r *Rarbg) Name() string {
	return name
}

func (r *Rarbg) Search(s string) ([]torrent.ProviderData, error) {
	res, err := r.api.Format(format).SearchString(s).Search()

	return data(res), err
}

func (r *Rarbg) SearchIMDB(id string) ([]torrent.ProviderData, error) {
	res, err := r.api.Format(format).SearchIMDb(id).Search()

	return data(res), err
}

func (r *Rarbg) SearchMediaType(mt torrent.MediaType, s string) ([]torrent.ProviderData, error) {

	var res torrentapi.TorrentResults
	var err error

	switch mt {
	case torrent.MediaTypeUnknown:
		return empty(), nil
	case torrent.MediaTypeMovie:
		res, err = r.api.Format(format).Movies().SearchString(s).Search()
	case torrent.MediaTypeTVShow:
		res, err = r.api.Format(format).TV().SearchString(s).Search()
	}

	return data(res), err
}

func (r *Rarbg) New() ([]torrent.ProviderData, error) {
	res, err := r.api.Format(format).List()

	return data(res), err
}

func (r *Rarbg) NewMediaType(mt torrent.MediaType) ([]torrent.ProviderData, error) {
	var res torrentapi.TorrentResults
	var err error

	switch mt {
	case torrent.MediaTypeUnknown:
		return empty(), nil
	case torrent.MediaTypeMovie:
		res, err = r.api.Format(format).Movies().List()
	case torrent.MediaTypeTVShow:
		res, err = r.api.Format(format).TV().List()
	}

	return data(res), err
}
