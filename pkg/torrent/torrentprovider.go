package torrent

type MediaType int

const (
	MediaTypeMovie MediaType = iota
	MediaTypeTVShow
	MediaTypeUnknown
)

// TorrentProvider knows how to communicate with a torrent indexer and
// retreive data from there
type TorrentProvider interface {
	Name() string                                              // The name of the provider
	SearchIMDB(string) ([]ProviderData, error)                 // Search for media matching the provided IMDB id
	Search(string) ([]ProviderData, error)                     // Search for torrents
	SearchMediaType(MediaType, string) ([]ProviderData, error) // Search for torrents of specified media type
	New() ([]ProviderData, error)                              // New torrents
	NewMediaType(MediaType) ([]ProviderData, error)            // New torrents of specified medya type
}

// ProviderData represents data comming from the provider
type ProviderData struct {
	Title        string    // String to help identify the torrents contents
	Type         MediaType // What kind of media is this
	Magnet       string    // The magnet link associated with this torrent
	Size         uint64    // Size in bytes of the torrent's contents
	IMDB         string    // The IMDB ID associated with this torrent
	ProviderName string    // Name of this torrent's provider
}

func (m MediaType) String() string {
	switch m {
	case MediaTypeMovie:
		return "movie"
	case MediaTypeTVShow:
		return "tv"
	default:
		return "unknown"
	}
}
