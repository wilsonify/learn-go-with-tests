package inputoutput

import (
	"io"
)

// FileSystemPlayerStore stores players in the filesystem.
type FileSystemPlayerStore struct {
	DatabaseSeeker io.ReadSeeker
}

// GetLeague returns the scores of all the players.
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.DatabaseSeeker.Seek(0, 0)
	league, _ := NewLeague(f.DatabaseSeeker)
	return league
}
