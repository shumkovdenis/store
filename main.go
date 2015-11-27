package main

import (
	"github.com/shumkovdenis/store/api"
	"github.com/shumkovdenis/store/database/bolt"
	"github.com/shumkovdenis/store/services/segmentation"
	"github.com/shumkovdenis/store/services/store"
	"log"
)

func main() {
	db, err := bolt.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	serviceStore := store.NewService(db)
	serviceSegmentation := segmentation.NewService(db)

	api.NewAPI(serviceStore, serviceSegmentation)
}
