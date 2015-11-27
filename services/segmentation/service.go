package segmentation

import (
	"github.com/shumkovdenis/store/database"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	return &Service{db}
}

func (service *Service) GetSegments(event, property string, from, to int64) (map[string]int, error) {
	tracks, err := service.db.GetTracks(event, from, to)
	if err != nil {
		return nil, err
	}

	segments := map[string]int{}
	for _, track := range tracks {
		if val, ok := track.Properties[property]; track.Event == event && ok {
			segments[val] += 1
		}
	}

	return segments, nil
}
