package middleware

import (
	"insider/src/infra/logger"
	"insider/src/infra/myerror"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

var excludedLogPaths = []string{
	"/api/v1/healthcheck",
}

func CreateLogMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogURI:     true,
			LogStatus:  true,
			LogLatency: true,
			LogMethod:  true,
			LogError:   true,
			BeforeNextFunc: func(c echo.Context) {
				setZLogger(c)
				logCallStarted(c)
			},
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logCallFinished(c, v)
				return nil
			},
		})
}

func setZLogger(c echo.Context) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	zLogger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("traceId", uuid.New().String()).
		Logger()

	c.SetLogger(logger.NewEchoZLogger(&zLogger))
}

func logCallStarted(c echo.Context) {

	if slices.Contains(excludedLogPaths, c.Path()) {
		return
	}

	zLogger := c.Logger().(logger.EchoZLogger).GetZLogger()

	zLogger.Info().
		Str("url", c.Path()).
		Str("method", c.Request().Method).
		Msg("Api call started")
}

func logCallFinished(c echo.Context, v middleware.RequestLoggerValues) {

	if slices.Contains(excludedLogPaths, c.Path()) {
		return
	}

	zLogger := c.Logger().(logger.EchoZLogger).GetZLogger()

	if v.Error != nil {
		status := findStatus(v.Error)
		code := findCode(v.Error)

		zEvent := zLogger.Error()

		if status >= 400 && status < 500 {
			zEvent = zLogger.Warn()
		}

		zEvent.Err(v.Error).
			Str("url", v.URI).
			Str("method", v.Method).
			Int("status", v.Status).
			Int("error_code", code).
			Dur("duration", v.Latency).
			Msg("Api call failed")

		return
	}

	zLogger.Info().
		Str("url", v.URI).
		Str("method", v.Method).
		Int("status", v.Status).
		Dur("duration", v.Latency).
		Msg("Api call succeeded")
}

func findStatus(err error) int {
	switch e := err.(type) {
	case *echo.HTTPError:
		return e.Code
	case *myerror.Error:
		return e.HttpCode
	default:
		return http.StatusInternalServerError
	}
}

func findCode(err error) int {
	switch e := err.(type) {
	case *myerror.Error:
		return e.ErrorCode
	default:
		return 1000
	}
}
