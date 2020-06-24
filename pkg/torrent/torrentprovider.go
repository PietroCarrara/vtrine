package torrent

type MediaType string

const (
	MediaTypeMovie   = "movie"
	MediaTypeTVShow  = "show"
	MediaTypeUnknown = "unknown"
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
