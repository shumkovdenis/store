package main

type Array struct {
	data []*Track
}

func NewArray() *Array {
	return &Array{[]*Track{}}
}

func (array *Array) Save(track *Track) {
	array.data = append(array.data, track)
}

func (array *Array) Count(event string, properties map[string]string) int {
	var c int
	for _, track := range array.data {
		f := true
		for k, _ := range properties {
			if _, e := track.Properties[k]; !e {
				f = false
				break
			}
		}
		if f {
			c++
		}
	}
	return c
}
