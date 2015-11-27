package store

import (
	"fmt"
	"github.com/shumkovdenis/store/database"
	"github.com/shumkovdenis/store/models"
	"strconv"
	"time"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	return &Service{db}
}

func (s *Service) SaveTrack(track *models.Track) error {
	if err := processTime(track); err != nil {
		return err
	}
	return s.db.SaveTrack(track)
}

func processTime(track *models.Track) error {
	if track.Properties == nil {
		track.Properties = map[string]string{}
	}
	if val, ok := track.Properties[models.PropertyTime]; !ok {
		timestamp := time.Now().Unix()
		track.Time = timestamp
		track.Properties[models.PropertyTime] = fmt.Sprintf("%d", timestamp)
	} else {
		timestamp, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		track.Time = timestamp
	}
	return nil
}
