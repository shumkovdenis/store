package statistics

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

func (s *Service) Segmentation(event, property string, from, to int64) ([]*models.Segment, error) {
	s.db.GetTracks(event, from, to)
	return nil, nil
}
