package services

type Segmentation interface {
	GetSegments(event, property string, from, to int64) (map[string]int, error)
}
