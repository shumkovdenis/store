package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shumkovdenis/store/models"
	"github.com/shumkovdenis/store/services"
)

type API struct {
	serviceStore        services.Store
	serviceSegmentation services.Segmentation
}

func NewAPI(serviceStore services.Store, serviceSegmentation services.Segmentation) *API {
	api := &API{serviceStore, serviceSegmentation}

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

	if err := api.serviceStore.SaveTrack(track); err != nil {
		return err
	}

	return c.String(http.StatusOK, "0")
}

func (api *API) segmentation(c *echo.Context) error {
	event := c.Query("event")
	property := c.Query("property")
	sFrom := c.Query("from")
	sTo := c.Query("to")

	iFrom, err := strconv.ParseInt(sFrom, 10, 64)
	if err != nil {
		return err
	}

	iTo, err := strconv.ParseInt(sTo, 10, 64)
	if err != nil {
		return err
	}

	segments, err := api.serviceSegmentation.GetSegments(event, property, iFrom, iTo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, segments)
}
