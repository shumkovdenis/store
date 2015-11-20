package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shumkovdenis/store/database/memory"
	"github.com/shumkovdenis/store/models"
	"github.com/shumkovdenis/store/services/store"
	"net/http"
)

type API struct {
	store *store.Service
}

func NewAPI() *API {
	db := memory.NewDB()

	service := store.NewService(db)

	api := &API{service}

	e := echo.New()

	e.Use(middleware.Logger())

	e.Post("/track", api.track)

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
