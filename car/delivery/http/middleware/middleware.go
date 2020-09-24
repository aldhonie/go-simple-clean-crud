package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// ResponseMessage Middleware
type ResponseMessage struct {
	Message string `json:"message"`
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// AuthHeader will handle request from middleware
func (m *GoMiddleware) AuthHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bool := strings.HasPrefix(c.Request().Header.Get("X-Api-Key"), viper.GetString(`middleware.x-api-key`))
		if bool != true {
			return c.JSON(404, ResponseMessage{Message: "Access Denied"})
		}
		return next(c)
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
