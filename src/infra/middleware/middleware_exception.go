package middleware

import (
	"insider/src/infra/api"
	"insider/src/infra/myerror"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ExceptionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, api.NewErrorResponse(err, "Record not found", 4040))
			}
			switch e := err.(type) {
			case *myerror.Error:
				return c.JSON(e.HttpCode, api.NewErrorResponse(e, e.Message, e.ErrorCode))
			case *echo.HTTPError:
				return c.JSON(e.Code, api.NewErrorResponse(e, e.Message.(string), e.Code))
			default:
				return c.JSON(http.StatusInternalServerError, api.NewErrorResponse(e, e.Error(), 1000))
			}
		}

		return nil
	}
}
