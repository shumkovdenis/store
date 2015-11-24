package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shumkovdenis/store/database/bolt"
	"github.com/shumkovdenis/store/models"
	"github.com/shumkovdenis/store/services/statistics"
	"github.com/shumkovdenis/store/services/store"
	"log"
	"net/http"
)

type API struct {
	store *store.Service
	stats *statistics.Service
}

func NewAPI() *API {
	db, err := bolt.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	store := store.NewService(db)
	stats := statistics.NewService(db)

	api := &API{store, stats}

	e := echo.New()

	e.Use(middleware.Logger())

	e.Post("/track", api.track)
	e.Get("/segmentation", api.segmentation)

	e.Run(":8080")

	return api
}

func (api *API) track(c *echo.Context) error {
	track := &models.Track{}

	if err := c.Bind(track); err != nil {
		return err
	}

	if err := api.store.Save(track); err != nil {
		return err
	}

	return c.String(http.StatusOK, "0")
}

func (api *API) segmentation(c *echo.Context) error {
	segments, err := api.stats.Segmentation("test", "page", 0, 1449361486)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, segments)
}
