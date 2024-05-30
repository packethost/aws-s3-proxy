package s3

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/packethost/aws-s3-proxy/internal/config"
)

// Health returns a handler function that returns a HTTP 200 response every time
func Health() echo.HandlerFunc {
	return echo.HandlerFunc(func(e echo.Context) error {
		h := config.Cfg.HTTPOpts
		res := e.Response()

		// Facility Header if set
		if h.Facility != "" {
			res.Header().Add("Facility", h.Facility)
		}

		res.WriteHeader(http.StatusOK)

		return nil
	})
}
