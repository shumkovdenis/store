package services

import "github.com/shumkovdenis/store/models"

type Store interface {
	SaveTrack(track *models.Track) error
}
