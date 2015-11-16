package main

import (
	"fmt"
	"math/rand"
)

var Tracks []*Track

func init() {
	for i := 0; i < 100000; i++ {
		Tracks = append(Tracks, GenerateTrack())
	}
}

func GenerateTrack() *Track {
	return &Track{0, "track", generateProperties(2)}
}

func generateProperties(count int) map[string]string {
	prop := map[string]string{}
	for i := 0; i < count; i++ {
		k := fmt.Sprintf("k%d", rand.Intn(count))
		v := fmt.Sprintf("v%d", rand.Intn(count))
		prop[k] = v
	}
	return prop
}
