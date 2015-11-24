package bolt

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(track.Event))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(track)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(track.Properties[models.PropertyTime]), encoded)
	})
}

func (db *DB) GetTracks(event string, from, to int64) ([]*models.Track, error) {
	db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(event)).Cursor()

		min := []byte(parseTime(from))
		max := []byte(parseTime(to))

		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}

		return nil
	})
	return nil, nil
}

func parseTime(value int64) string {
	tm := time.Unix(value, 0)
	return tm.Format(time.RFC3339)
}
