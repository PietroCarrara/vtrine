package torrent

// TorrentClient kows how to communicate with a torrent downloader and
// interact with it
type TorrentClient interface {
	DownloadMovie(string) error // Add a magnet torrent of a movie to the download queue
	DownloadShow(string) error  // Add a magnet torrent of a tv show to the download queue
	DownloadAnime(string) error // Add a magnet torrent of a anime to the download queue
}
