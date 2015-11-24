package memory

import "github.com/shumkovdenis/store/models"

type DB struct {
	data []*models.Track
}

func NewDB() (*DB, error) {
	data := []*models.Track{}

	db := &DB{data}

	return db, nil
}

func (db *DB) Save(track *models.Track) error {
	db.data = append(db.data, track)

	return nil
}
