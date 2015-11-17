package main
import "fmt"

func main() {
	var store Store = NewArray()
	FillGenerate(store)
	c := store.Count("track", map[string]string{"k1": "v1"})
	fmt.Println(c)
}
