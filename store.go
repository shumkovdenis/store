package main

type Track struct {
	Time       int64
	Event      string
	Properties map[string]string
}

type Store interface {
	Save(event *Track)
	Count(event string, properties map[string]string) int
}

func FillGenerate(store Store, count int) {
	for i := 0; i < count; i++ {
		track := GenerateTrack()
		store.Save(track)
	}
}

func FillExist(store Store, count int) {
	for _, track := range Tracks {
		store.Save(track)
	}
}
