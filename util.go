package main

import (
	"fmt"
	"math/rand"
)

const COUNT_TRACKS int = 1000000
const COUNT_PROPERTIES int = 2

var Tracks []*Track

func init() {
	for i := 0; i < COUNT_TRACKS; i++ {
		Tracks = append(Tracks, generateTrack(COUNT_PROPERTIES))
	}
}

func generateProperties(countProperties int) map[string]string {
	properties := map[string]string{}
	for i := 0; i < countProperties; i++ {
		k := fmt.Sprintf("k%d", 1+rand.Intn(countProperties))
		v := fmt.Sprintf("v%d", 1+rand.Intn(countProperties))
		properties[k] = v
	}
	return properties
}

func generateTrack(countProperties int) *Track {
	return &Track{0, "track", generateProperties(1 + rand.Intn(countProperties))}
}

func FillGenerate(store Store) {
	for i := 0; i < COUNT_TRACKS; i++ {
		track := generateTrack(COUNT_PROPERTIES)
		store.Save(track)
	}
}

func FillExist(store Store) {
	for _, track := range Tracks {
		store.Save(track)
	}
}
