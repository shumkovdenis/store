package bolt

import (
	"bytes"
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/shumkovdenis/store/models"
	"time"
)

type DB struct {
	*bolt.DB
}

func NewDB() (*DB, error) {
	db, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) SaveTrack(track *models.Track) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(track.Event))
		if err != nil {
			return err
		}

		key := []byte(formatTime(track.Time))
		value, err := json.Marshal(track)
		if err != nil {
			return err
		}

		return bucket.Put(key, value)
	})

	return err
}

func (db *DB) GetTracks(event string, from, to int64) ([]*models.Track, error) {
	var tracks []*models.Track
	err := db.View(func(tx *bolt.Tx) error {
		cursor := tx.Bucket([]byte(event)).Cursor()

		min := []byte(formatTime(from))
		max := []byte(formatTime(to))

		for key, value := cursor.Seek(min); key != nil && bytes.Compare(key, max) <= 0; key, value = cursor.Next() {
			track := &models.Track{}
			if err := json.Unmarshal(value, track); err != nil {
				return err
			}
			tracks = append(tracks, track)
		}

		return nil
	})

	return tracks, err
}

func formatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(time.RFC3339)
}
