package torrent

// TorrentClient kows how to communicate with a torrent downloader and
// interact with it
type TorrentClient interface {
	DownloadMovie(data ProviderData) error // Add a torrent of a movie to the download queue
	DownloadShow(data ProviderData) error  // Add a torrent of a tv show to the download queue
	DownloadAnime(data ProviderData) error // Add a torrent of a anime to the download queue

	ListMovies() ([]ClientData, error) // List all movies
	ListShows() ([]ClientData, error)  // List all tv shows
	ListAnimes() ([]ClientData, error) // List all animes

	GetFreeSpace() (uint64, error) // Returns the free space in the disk, in bytes
	GetUsedSpace() (uint64, error) // Returns the disk size, in bytes

	RemoveTorrent(id string) error // Remove a torrent
}

type ClientData struct {
	ID       string  // Internal ID
	Title    string  // Name of the torrent
	Size     uint64  // Size in bytes
	IMDB     string  // IMDB id of this media
	Complete bool    // Is the download finished?
	Paused   bool    // Is this torrent currently paused?
	Progress float32 // Download progress in range [0, 1]
}
