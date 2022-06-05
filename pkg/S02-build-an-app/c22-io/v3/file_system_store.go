package inputoutput

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore stores players in the filesystem.
type FileSystemPlayerStore struct {
	DatabaseSeeker io.ReadWriteSeeker
}

// GetLeague returns the scores of all the players.
func (f *FileSystemPlayerStore) GetLeague() League {
	f.DatabaseSeeker.Seek(0, 0)
	league, _ := NewLeague(f.DatabaseSeeker)
	return league
}

// GetPlayerScore retrieves a player's score.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin will store a win for a player, incrementing wins if already known.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	f.DatabaseSeeker.Seek(0, 0)
	json.NewEncoder(f.DatabaseSeeker).Encode(league)
}
