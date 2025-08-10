package middleware

import (
	"insider/src/infra/myerror"

	"github.com/labstack/echo/v4"
)

func RecoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (returnErr error) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = myerror.NewInternalServerError(err, "Errorless panic how is that possible?", 1000)
				}
				returnErr = err
			}
		}()

		return next(c)
	}
}
