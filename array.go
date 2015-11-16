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
		if track.Event == event {
			e := true
			for k, _ := range properties {
				_, f := track.Properties[k]
				if !f {
					e = false
					break
				}
			}
			if e {
				c++
			}
		}
	}
	return c
}
