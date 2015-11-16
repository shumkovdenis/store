package main

import "testing"

func BenchmarkArraySave1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var store Store = NewArray()
		FillGenerate(store, 100000)
	}
}

func BenchmarkArraySave2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var store Store = NewArray()
		FillExist(store, 100000)
	}
}

//func BenchmarkArrayCount(b *testing.B) {
//	var store Store = NewArray()
//	Fill(store, 100000)
//	b.Log("begin")
//	for n := 0; n < b.N; n++ {
//		c := store.Count("track", map[string]string{"k1": "v1"})
//		b.Log("get", c)
//	}
//}
