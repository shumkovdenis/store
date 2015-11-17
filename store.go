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
