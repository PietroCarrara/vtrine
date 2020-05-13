package rarbg

import (
	"github.com/PietroCarrara/vtrine/pkg/torrentprovider"
	"github.com/qopher/go-torrentapi"
)

const format = "json_extended"

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
	return "rarbg"
}

func (r *Rarbg) Search(s string) ([]torrentprovider.TorrentData, error) {
	res, err := r.api.Format(format).SearchString(s).Search()

	return data(res), err
}

func (r *Rarbg) SearchMediaType(mt torrentprovider.MediaType, s string) ([]torrentprovider.TorrentData, error) {

	var res torrentapi.TorrentResults
	var err error

	switch mt {
	case torrentprovider.MediaTypeAnime, torrentprovider.MediaTypeUnknown:
		return empty(), nil
	case torrentprovider.MediaTypeMovie:
		res, err = r.api.Format(format).Movies().SearchString(s).Search()
	case torrentprovider.MediaTypeTVShow:
		res, err = r.api.Format(format).TV().SearchString(s).Search()
	}

	return data(res), err
}

func (r *Rarbg) New() ([]torrentprovider.TorrentData, error) {
	res, err := r.api.Format(format).List()

	return data(res), err
}

func (r *Rarbg) NewMediaType(mt torrentprovider.MediaType) ([]torrentprovider.TorrentData, error) {
	var res torrentapi.TorrentResults
	var err error

	switch mt {
	case torrentprovider.MediaTypeAnime, torrentprovider.MediaTypeUnknown:
		return empty(), nil
	case torrentprovider.MediaTypeMovie:
		res, err = r.api.Format(format).Movies().List()
	case torrentprovider.MediaTypeTVShow:
		res, err = r.api.Format(format).TV().List()
	}

	return data(res), err
}
