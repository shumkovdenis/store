package store

import (
	"github.com/shumkovdenis/store/database"
	"github.com/shumkovdenis/store/models"
	"strconv"
	"time"
	"fmt"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	service := &Service{db}

	return service
}

func (s *Service) Save(track *models.Track) error {
	if track.Properties == nil {
		track.Properties = map[string]string{}
	}

	if err := processTime(track.Properties); err != nil {
		return err
	}
	fmt.Println(track.Properties)
	return s.db.SaveTrack(track)
}

func processTime(properties map[string]string) error {
	tm := time.Now()
	if value, ok := properties[models.PropertyTime]; ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		tm = time.Unix(i, 0)
	}
	properties[models.PropertyTime] = tm.Format(time.RFC3339)
	return nil
}
