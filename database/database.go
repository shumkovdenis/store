package database

import "github.com/shumkovdenis/store/models"

type DB interface {
	SaveTrack(track *models.Track) error
	GetTracks(event string, from, to int64) ([]*models.Track, error)
}
