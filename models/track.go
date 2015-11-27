package models

const (
	PropertyTime = "time"
)

type Track struct {
	Time       int64
	Event      string            `json:"event"`
	Properties map[string]string `json:"properties"`
}
