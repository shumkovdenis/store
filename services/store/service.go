package store

import (
	"github.com/shumkovdenis/store/database"
	"github.com/shumkovdenis/store/models"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	service := &Service{db}

	return service
}

func (s *Service) Save(track *models.Track) error {
	return s.db.Save(track)
}
