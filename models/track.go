package models

const (
	PropertyTime = "time"
)

type Track struct {
	Event      string            `json:"event"`
	Properties map[string]string `json:"properties"`
}
