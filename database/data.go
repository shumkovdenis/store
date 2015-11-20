package database

import "github.com/shumkovdenis/store/models"

type DB interface {
	Save(track *models.Track) error
}
